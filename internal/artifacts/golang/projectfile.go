package golang

import "github.com/peterj/create-k8s-app/internal/artifacts"

// NewGolangProjectFile creates a new project file for Golang
func NewGolangProjectFile(templateName, targetFileName string) *artifacts.ProjectFile {
	return artifacts.NewProjectFile(templateName, targetFileName, "golang")
}
