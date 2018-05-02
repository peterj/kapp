package artifacts

// ProjectInfo holds info we need to create a project
type ProjectInfo struct {
	ApplicationName  string
	PackageName      string
	VersionFileName  string
	DockerRepository string
}

// NewProjectInfo creates a new ProjectInfo
func NewProjectInfo(applicationName, packageName, dockerRepository string) *ProjectInfo {
	return &ProjectInfo{
		ApplicationName:  applicationName,
		PackageName:      packageName,
		DockerRepository: dockerRepository,
		VersionFileName:  "VERSION.txt",
	}
}
