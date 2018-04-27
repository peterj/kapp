package artifacts

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/peterj/create-k8s-app/internal/utils"
	"github.com/pkg/errors"
)

// ProjectFile represents a file that's part of a project
type ProjectFile struct {
	Language           string
	Template           string
	TargetFileNamePath string
}

// NewProjectFile creates a new project file
func NewProjectFile(template, targetFileName, language string) *ProjectFile {
	return &ProjectFile{
		Template:           template,
		TargetFileNamePath: targetFileName,
		Language:           language,
	}
}

// Write writes contents to the project file
func (p *ProjectFile) Write(rootPath string, contents []byte) error {
	fullPath := path.Join(rootPath, p.TargetFileNamePath)
	if err := utils.CreateFolder(fullPath); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Write: %s", fullPath))
	}
	if err := ioutil.WriteFile(fullPath, contents, 0644); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Write: %s", fullPath))
	}
	return nil
}
