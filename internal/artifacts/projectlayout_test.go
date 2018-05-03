package artifacts_test

import (
	"testing"

	. "github.com/peterj/kapp/internal/artifacts"
)

func TestProjectLayout(t *testing.T) {
	tt := []struct {
		testName     string
		projectInfo  *ProjectInfo
		projectFiles []*ProjectFile
	}{
		{
			testName: "create ProjectLayout instance",
			projectInfo: &ProjectInfo{
				ApplicationName:  "app",
				PackageName:      "package",
				VersionFileName:  "VERSION.txt",
				DockerRepository: "dockerrepo",
			},
			projectFiles: []*ProjectFile{
				{
					Delimiters: &Delims{
						Left:  "{{",
						Right: "}}",
					},
					TargetFileNamePath: "firstfile",
					Template:           "template",
				},
				{
					Delimiters: &Delims{
						Left:  "{{",
						Right: "}}",
					},
					TargetFileNamePath: "secondfile",
					Template:           "template",
				},
			},
		},
	}

	t.Log("Given the need to test project layout creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				creator, _ := NewProjectLayout(tst.projectInfo, tst.projectFiles)
				layout, _ := creator.(*ProjectLayout)

				if len(layout.Files) != len(tst.projectFiles) {
					t.Fatalf("\t%s\tShould have the correct number of files : exp[%d] got[%d]\n", failed, len(tst.projectFiles), len(layout.Files))
				}
				t.Logf("\t%s\tShould have the correct number of files\n", succeeded)

				if layout.ProjectInfo.ApplicationName != tst.projectInfo.ApplicationName {
					t.Fatalf("\t%s\tShould have the correct app name : exp[%s] got[%s]\n", failed, tst.projectInfo.ApplicationName, layout.ProjectInfo.ApplicationName)
				}
				t.Logf("\t%s\tShould have the correct app name\n", succeeded)

				if layout.ProjectInfo.DockerRepository != tst.projectInfo.DockerRepository {
					t.Fatalf("\t%s\tShould have the correct docker repo name : exp[%s] got[%s]\n", failed, tst.projectInfo.DockerRepository, layout.ProjectInfo.DockerRepository)
				}
				t.Logf("\t%s\tShould have the correct docker repo name\n", succeeded)

				if layout.ProjectInfo.PackageName != tst.projectInfo.PackageName {
					t.Fatalf("\t%s\tShould have the correct package name : exp[%s] got[%s]\n", failed, tst.projectInfo.PackageName, layout.ProjectInfo.PackageName)
				}
				t.Logf("\t%s\tShould have the correct package name\n", succeeded)

				if layout.ProjectInfo.VersionFileName != tst.projectInfo.VersionFileName {
					t.Fatalf("\t%s\tShould have the correct version file name : exp[%s] got[%s]\n", failed, tst.projectInfo.VersionFileName, layout.ProjectInfo.VersionFileName)
				}
				t.Logf("\t%s\tShould have the correct version file name\n", succeeded)
			}
		}
	}
}
