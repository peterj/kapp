package golang

import (
	"bytes"
	"html/template"

	"github.com/peterj/create-k8s-app/internal/artifacts"
	"github.com/pkg/errors"
)

const (
	// DefaultVersionFileName is a name of the file that holds the version number
	DefaultVersionFileName = "VERSION.txt"
	templateName           = "Makefile.templ"
	// languageName needs to also match the folder name under /templates
	languageName = "golang"
)

// Makefile represents a Makefile for Go project
type Makefile struct {
	ApplicationName string
	VersionFileName string
	PackageName     string
	// TODO: Add Go architectures here
}

// NewMakeFile creates a new Makefile
func NewMakeFile(applicationName, packageName string) *Makefile {
	return &Makefile{
		ApplicationName: applicationName,
		PackageName:     packageName,
		VersionFileName: DefaultVersionFileName,
	}
}

// TemplateFileInfo gets information about the template
func (g *Makefile) TemplateFileInfo() *artifacts.TemplateFileInfo {
	return artifacts.NewTemplateFileInfo(templateName, languageName)
}

// Create generates an artifact from a template file
func (g *Makefile) Create() ([]byte, error) {
	templateInfo := g.TemplateFileInfo()
	fullPath, err := templateInfo.FullPath()
	if err != nil {
		return []byte{}, errors.Wrap(err, "Create")
	}

	t := template.Must(template.New(templateInfo.Name).ParseFiles(fullPath))

	var result bytes.Buffer
	if err = t.Execute(&result, g); err != nil {
		return []byte{}, errors.Wrap(err, "Create")
	}
	return result.Bytes(), nil
}
