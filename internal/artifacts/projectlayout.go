package artifacts

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/peterj/kapp/internal/utils"
	"github.com/pkg/errors"
)

// ProjectLayout holds the information about the project and files
// in that project
type ProjectLayout struct {
	ProjectInfo *ProjectInfo
	Files       []*ProjectFile
}

// NewProjectLayout creates a new project layout using project and project files
func NewProjectLayout(project *ProjectInfo, projectFiles []*ProjectFile) (LayoutCreator, error) {
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
		return errors.Wrap(err, fmt.Sprintf("Create: output folder: '%s'", outputFolder))
	}

	for _, f := range p.Files {
		tmpl := template.New("tmpl").Delims(f.Delimiters.Left, f.Delimiters.Right)
		tmpl, err := tmpl.Parse(f.Template)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Create: parse template: '%s'", f.TargetFileNamePath))
		}

		var result bytes.Buffer
		if err := tmpl.ExecuteTemplate(&result, "tmpl", p.ProjectInfo); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Create: execute template: '%s'", result.String()))
		}
		if err := f.Write(outputPath, result.Bytes()); err != nil {
			return errors.Wrap(err, "Create: write")
		}
	}
	return nil
}
