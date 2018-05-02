package artifacts_test

import (
	"testing"

	. "github.com/peterj/kapp/internal/artifacts"
)

func TestProjectInfo(t *testing.T) {
	tt := []struct {
		testName             string
		appName              string
		versionFileName      string
		packageName          string
		dockerRepositoryName string
	}{
		{
			testName:             "create ProjectInfo instance",
			appName:              "myapp",
			versionFileName:      "VERSION.txt",
			packageName:          "github.com/test/myapp",
			dockerRepositoryName: "myrepo.registry.io/something",
		},
	}

	t.Log("Give the need to test project info creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				projectInfo := NewProjectInfo(tst.appName, tst.packageName, tst.dockerRepositoryName)

				if projectInfo.ApplicationName != tst.appName {
					t.Fatalf("\t%s\tShould have the correct app name : exp[%s] got[%s]\n", failed, tst.appName, projectInfo.ApplicationName)
				}
				t.Logf("\t%s\tShould have the correct app name\n", succeeded)

				if projectInfo.PackageName != tst.packageName {
					t.Fatalf("\t%s\tShould have the correct package name : exp[%s] got[%s]\n", failed, tst.packageName, projectInfo.PackageName)
				}
				t.Logf("\t%s\tShould have the correct package name\n", succeeded)

				if projectInfo.DockerRepository != tst.dockerRepositoryName {
					t.Fatalf("\t%s\tShould have the correct Docker repository name : exp[%s] got[%s]\n", failed, tst.dockerRepositoryName, projectInfo.DockerRepository)
				}
				t.Logf("\t%s\tShould have the correct Docker repository name\n", succeeded)

				if projectInfo.VersionFileName != tst.versionFileName {
					t.Fatalf("\t%s\tShould have the correct version file name : exp[%s] got[%s]\n", failed, tst.versionFileName, projectInfo.VersionFileName)
				}
				t.Logf("\t%s\tShould have the correct version file name\n", succeeded)
			}
		}
	}
}
