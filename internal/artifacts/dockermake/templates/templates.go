package templates

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
	kubectl create ns $(3)
	helm install $(1) $(2) --namespace $(3) --set=image.repository=$(4)
endef

# Upgrades an existing Helm chart
define helm_upgrade
	helm upgrade $(1) $(2) --namespace $(3) --set=image.repository=$(4)
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
	$(call helm_install,$(RELEASE_NAME),$(HELM_CHART_NAME),$(KUBE_NAMESPACE),$(REGISTRY_NAME)/$(IMAGE_NAME))

.PHONY: upgrade.app
upgrade.app:
	@echo "-> $@"
	$(call helm_upgrade,$(RELEASE_NAME),$(HELM_CHART_NAME),$(KUBE_NAMESPACE),$(REGISTRY_NAME)/$(IMAGE_NAME))

.PHONY: upgrade
upgrade: build.image push.image upgrade.app
`
