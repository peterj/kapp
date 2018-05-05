package dockermake

import (
	"github.com/peterj/kapp/internal/artifacts"
	dockertemplates "github.com/peterj/kapp/internal/artifacts/dockermake/templates"
)

// ProjectFiles returns docker.mk project file
func ProjectFiles() []*artifacts.ProjectFile {
	return []*artifacts.ProjectFile{
		artifacts.NewProjectFile(dockertemplates.DockerMkFile, "docker.mk", artifacts.NewDefaultDelims()),
	}
}
