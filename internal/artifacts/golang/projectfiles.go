package golang

import (
	"github.com/peterj/kapp/internal/artifacts"
	gotemplates "github.com/peterj/kapp/internal/artifacts/golang/templates"
)

// ProjectFiles returns a slice of files needed for a Go project
func ProjectFiles() []*artifacts.ProjectFile {
	goTemplateDelimiters := artifacts.NewDefaultDelims()

	return []*artifacts.ProjectFile{
		artifacts.NewProjectFile(gotemplates.Makefile, "Makefile", goTemplateDelimiters),
		artifacts.NewProjectFile(gotemplates.VersionTxtFile, "VERSION.txt", goTemplateDelimiters),
		artifacts.NewProjectFile(gotemplates.Dockerfile, "Dockerfile", goTemplateDelimiters),
		artifacts.NewProjectFile(gotemplates.VersionGo, "version/version.go", goTemplateDelimiters),
		artifacts.NewProjectFile(gotemplates.DockerMkFile, "docker.mk", goTemplateDelimiters),
		artifacts.NewProjectFile(gotemplates.MainGo, "main.go", goTemplateDelimiters),
		artifacts.NewProjectFile(gotemplates.GitIgnore, ".gitignore", goTemplateDelimiters),
	}
}
