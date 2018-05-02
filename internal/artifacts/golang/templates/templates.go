package template

// VersionTxtFile holds the contents of the VERSION.txt
const VersionTxtFile = `0.1.0`

// Makefile holds the contents of the Makefile
const Makefile = `include docker.mk 
NAME:={{ .ApplicationName }}
PKG:={{ .PackageName }}
GOOSARCHES=darwin/amd64
VERSION_FILE:={{ .VersionFileName }}

VERSION=$(shell cat ./$(VERSION_FILE))
GITCOMMIT:=$(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES:=$(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif

# Sets the actual GITCOMMIT and VERSION values 
VERSION_INFO=-X $(PKG)/version.GITCOMMIT=$(GITCOMMIT) -X $(PKG)/version.VERSION=$(VERSION)

# Set the linker flags
GO_LDFLAGS=-ldflags "-w $(VERSION_INFO)"

all: build fmt lint test vet

# Builds the binary
.PHONY: build
build:
	@echo "-> $@"
	CGO_ENABLED=0 go build -i -installsuffix cgo ${GO_LDFLAGS} -o $(NAME) .

# Gofmts all code (sans vendor folder) just in case not using automatic formatting
.PHONY: fmt
fmt: 
	@echo "-> $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

# Runs golint
.PHONY: lint
lint:
	@echo "-> $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

# Runs all tests
.PHONY: test
test:
	@echo "-> $@"
	@go test -v $(shell go list ./... | grep -v vendor)

# Runs govet
.PHONY: vet
vet:
	@echo "-> $@"
	@go vet $(shell go list ./... | grep -v vendor) | tee /dev/stderr

# Bumps the version of the service
.PHONY: bump-version
BUMP := patch
bump-version:
	$(eval NEW_VERSION = $(shell curl https://bump.semver.xyz/$(BUMP)?version=$(VERSION)))
	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
	echo $(NEW_VERSION) > VERSION.txt
	git commit -vsam "Bump version to $(NEW_VERSION)"`

// Dockerfile holds the contents of the Dockerfile
const Dockerfile = `FROM golang:1.9.2 as builder

RUN mkdir -p /go/src/{{ .PackageName }}
WORKDIR  /go/src/{{ .PackageName }}

RUN go get -u github.com/golang/dep/cmd/dep 
COPY . .
RUN dep ensure -v
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir -p public
COPY --from=builder /go/src/{{ .PackageName }}/{{ .ApplicationName }} .
ENTRYPOINT [ "./{{ .ApplicationName }}" ]`

// VersionGo holds the contents of the version.go file
const VersionGo = `package version

// VERSION indicates the binary version
var VERSION string

// GITCOMMIT indicates the git hash binary was built off of
var GITCOMMIT string`

// DockerMkFile holds the contents of the docker.mk file
const DockerMkFile = `VERSION=$(shell cat ./{{ .VersionFileName }})

# Docker settings (make sure DOCKER_REGISTRY environment variable is set)
DOCKERFILE:=Dockerfile
IMAGE_NAME={{ .ApplicationName }}
REGISTRY_NAME:=$(DOCKER_REGISTRY)
FULL_IMAGE_NAME=$(REGISTRY_NAME)/$(IMAGE_NAME):$(VERSION)

# Kubernetes/Helm settings
KUBE_NAMESPACE:={{ .ApplicationName }}
RELEASE_NAME:={{ .ApplicationName }}
HELM_CHART_NAME:=helm/{{ .ApplicationName }}

# Builds a docker image
define build_image
	docker build -f $(1) -t $(2) .
endef

# Pushes a docker image to registry
define push_image
	docker push $(1)
endef

# Installs a new Helm chart
define helm_install
	helm install --name $(1) --namespace $(2) --set=image.repository=$(3) --set=image.tag=$(4) $(5)
endef

# Upgrades an existing Helm chart
define helm_upgrade
	helm upgrade $(1) --namespace $(2) --set=image.repository=$(3) --set=image.tag=$(4) $(5) --recreate-pods
endef

.PHONY: build.image
build.image:
	@echo "-> $@"
	$(call build_image, $(DOCKERFILE), $(FULL_IMAGE_NAME))

.PHONY: push.image
push.image:
	@echo "-> $@"
	$(call push_image, $(FULL_IMAGE_NAME))

.PHONE: install.app
install.app:
	@echo "-> $@"
	$(call helm_install,$(RELEASE_NAME),$(KUBE_NAMESPACE),$(REGISTRY_NAME)/$(IMAGE_NAME),$(VERSION),$(HELM_CHART_NAME))

.PHONY: upgrade.app
upgrade.app:
	@echo "-> $@"
	$(call helm_upgrade,$(RELEASE_NAME),$(KUBE_NAMESPACE),$(REGISTRY_NAME)/$(IMAGE_NAME),$(VERSION),$(HELM_CHART_NAME))

.PHONY: upgrade
upgrade: build.image push.image upgrade.app
`

// MainGo holds the contents of the Main.go file
const MainGo = `package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"{{ .PackageName }}/version"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello\n"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	port := ":8080"
	fmt.Printf("running on : %s\n", port)
	fmt.Printf("version    : %s\n", version.VERSION)
	fmt.Printf("git hash   : %s\n", version.GITCOMMIT)
	fmt.Printf("go version : %s\n", runtime.Version())
	fmt.Printf("go compiler: %s\n", runtime.Compiler)
	fmt.Printf("platform   : %s/%s\n", runtime.GOOS, runtime.GOARCH)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}`

// GitIgnore holds the contents of the .gitignore file
const GitIgnore = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out`
