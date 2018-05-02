package artifacts

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/peterj/kapp/internal/utils"
	"github.com/pkg/errors"
)

// ProjectFile represents a file that's part of a project
type ProjectFile struct {
	Template           string
	TargetFileNamePath string
	Delimiters         *Delims
}

// NewProjectFile creates a new project file
func NewProjectFile(template, targetFileName string, delims *Delims) *ProjectFile {
	return &ProjectFile{
		Template:           template,
		TargetFileNamePath: targetFileName,
		Delimiters:         delims,
	}
}

// Write writes contents to the project file
func (p *ProjectFile) Write(rootPath string, contents []byte) error {
	fullPath := path.Join(rootPath, p.TargetFileNamePath)
	if err := utils.CreateFolder(fullPath); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Write: create folder: '%s'", fullPath))
	}
	if err := ioutil.WriteFile(fullPath, contents, 0644); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Write: file: '%s'", fullPath))
	}
	return nil
}
