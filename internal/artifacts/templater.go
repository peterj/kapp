package artifacts

import (
	"os"
	"path"

	"github.com/pkg/errors"
)

// Templater is implemented by anything that is a templated file
type Templater interface {
	TemplateFileInfo() (*TemplateFileInfo, error)
}

// TemplateFileInfo holds the information about a template file
type TemplateFileInfo struct {
	Name     string
	Language string
}

// NewTemplateFileInfo creates a new template file info
func NewTemplateFileInfo(name, language string) *TemplateFileInfo {
	return &TemplateFileInfo{
		Name:     name,
		Language: language,
	}
}

// FullPath returns the full path to the template
func (t *TemplateFileInfo) FullPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "FullPath")
	}
	return path.Join(wd, "templates", t.Language, t.Name), nil
}
