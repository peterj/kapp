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
	TemplateName       string
	TargetFileNamePath string
}

// NewProjectFile creates a new project file
func NewProjectFile(templateName, targetFileName, language string) *ProjectFile {
	return &ProjectFile{
		TemplateName:       templateName,
		TargetFileNamePath: targetFileName,
		Language:           language,
	}
}

// GetTemplateFullPath gets the full path to the template
func (p *ProjectFile) GetTemplateFullPath() (string, error) {
	path, err := utils.FullPath("templates", p.Language, p.TemplateName)
	if err != nil {
		return "", errors.Wrap(err, "GetTemplateFullPath")
	}
	return path, nil
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
