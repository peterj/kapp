package helm

import (
	"path"

	"github.com/peterj/kapp/internal/artifacts"
	helmtemplates "github.com/peterj/kapp/internal/artifacts/helm/templates"
)

// ProjectFiles returns a slice of files need for a Helm chart
func ProjectFiles(applicationName string) []*artifacts.ProjectFile {
	// Since Helm uses "{{" and "}}" in its templates, we have to use [[, ]]
	helmDelimiters := artifacts.NewDelims("[[", "]]")
	rootFolder := path.Join("helm", applicationName)
	return []*artifacts.ProjectFile{
		artifacts.NewProjectFile(helmtemplates.HelpersTpl, path.Join(rootFolder, "templates", "_helpers.tpl"), helmDelimiters),
		artifacts.NewProjectFile(helmtemplates.DeploymentYaml, path.Join(rootFolder, "templates", "deployment.yaml"), helmDelimiters),
		artifacts.NewProjectFile(helmtemplates.ServiceYaml, path.Join(rootFolder, "templates", "service.yaml"), helmDelimiters),
		artifacts.NewProjectFile(helmtemplates.ChartYaml, path.Join(rootFolder, "Chart.yaml"), helmDelimiters),
		artifacts.NewProjectFile(helmtemplates.ValuesYaml, path.Join(rootFolder, "values.yaml"), helmDelimiters),
	}
}
