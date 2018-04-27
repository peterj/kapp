package golang

import (
	"bytes"
	"html/template"

	"github.com/peterj/create-k8s-app/internal/artifacts"
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

// NewProjectLayout creates a new project layout for a Go project
func NewProjectLayout(project *ProjectInfo) (artifacts.LayoutCreator, error) {
	projectFiles := []*artifacts.ProjectFile{
		NewGolangProjectFile("Makefile.templ", "Makefile"),
		NewGolangProjectFile("VERSION.txt.templ", "VERSION.txt"),
		NewGolangProjectFile("Dockerfile.templ", "Dockerfile"),
		NewGolangProjectFile("version.go.templ", "version/version.go"),
		NewGolangProjectFile("docker.mk.templ", "docker.mk"),
		NewGolangProjectFile("main.go.templ", "main.go"),
	}
	return &ProjectLayout{
		Files:       projectFiles,
		ProjectInfo: project,
	}, nil
}

// GetTemplatePaths gets an array of template paths for the
// project files in the layout
func (p *ProjectLayout) GetTemplatePaths() ([]string, error) {
	var paths []string
	for _, f := range p.Files {
		path, err := f.GetTemplateFullPath()
		if err != nil {
			return []string{}, errors.Wrap(err, "GetTemplatePaths")
		}
		paths = append(paths, path)
	}
	return paths, nil
}

// Create will write files in project layout to the output folder
func (p *ProjectLayout) Create(outputFolder string) error {
	projectFilePaths, err := p.GetTemplatePaths()
	if err != nil {
		return errors.Wrap(err, "Create")
	}

	// Get the output path by appending the output folder to
	// the current working folder
	outputPath, err := utils.FullPath(outputFolder)
	if err != nil {
		return errors.Wrap(err, "Create")
	}

	t := template.Must(template.New(p.Files[0].TemplateName).ParseFiles(projectFilePaths...))
	for _, f := range p.Files {
		var result bytes.Buffer
		if err := t.ExecuteTemplate(&result, f.TemplateName, p.ProjectInfo); err != nil {
			return errors.Wrap(err, "Create")
		}
		if err := f.Write(outputPath, result.Bytes()); err != nil {
			return errors.Wrap(err, "Create")
		}

	}
	return nil
}
