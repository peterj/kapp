package golang

import (
	"bytes"
	"html/template"

	"github.com/peterj/create-k8s-app/internal/artifacts"
	gotemplates "github.com/peterj/create-k8s-app/internal/artifacts/golang/templates"
	"github.com/peterj/create-k8s-app/internal/utils"
	"github.com/pkg/errors"
)

// ProjectLayout holds the information about the project and files
// in that project
type ProjectLayout struct {
	ProjectInfo *ProjectInfo
	Files       []*artifacts.ProjectFile
}

// ProjectInfo holds info we need to create a Go project
type ProjectInfo struct {
	ApplicationName string
	PackageName     string
	VersionFileName string
}

// NewGolangProjectFile creates a new project file for Golang
func NewGolangProjectFile(template, targetFileName string) *artifacts.ProjectFile {
	return artifacts.NewProjectFile(template, targetFileName, "golang")
}

// NewProjectLayout creates a new project layout for a Go project
func NewProjectLayout(project *ProjectInfo) (artifacts.LayoutCreator, error) {
	projectFiles := []*artifacts.ProjectFile{
		NewGolangProjectFile(gotemplates.Makefile, "Makefile"),
		NewGolangProjectFile(gotemplates.VersionTxtFile, "VERSION.txt"),
		NewGolangProjectFile(gotemplates.Dockerfile, "Dockerfile"),
		NewGolangProjectFile(gotemplates.VersionGo, "version/version.go"),
		NewGolangProjectFile(gotemplates.DockerMkFile, "docker.mk"),
		NewGolangProjectFile(gotemplates.MainGo, "main.go"),
	}
	return &ProjectLayout{
		Files:       projectFiles,
		ProjectInfo: project,
	}, nil
}

// Create will write files in project layout to the output folder
func (p *ProjectLayout) Create(outputFolder string) error {
	// Get the output path by appending the output folder to
	// the current working folder
	outputPath, err := utils.FullPath(outputFolder)
	if err != nil {
		return errors.Wrap(err, "Create")
	}

	for _, f := range p.Files {
		tmpl := template.New("tmpl")
		tmpl, err := tmpl.Parse(f.Template)
		if err != nil {
			return errors.Wrap(err, "Create")
		}

		var result bytes.Buffer
		if err := tmpl.ExecuteTemplate(&result, "tmpl", p.ProjectInfo); err != nil {
			return errors.Wrap(err, "Create")
		}
		if err := f.Write(outputPath, result.Bytes()); err != nil {
			return errors.Wrap(err, "Create")
		}
	}
	return nil
}
