package golang_test

import (
	"testing"

	"github.com/peterj/create-k8s-app/internal/artifacts/golang"
)

const succeeded = "\u2713"
const failed = "\u2717"

func TestMakefileCreation(t *testing.T) {
	tt := []struct {
		testName     string
		appName      string
		packageName  string
		languageName string
		templateName string
	}{
		{
			testName:     "create Go Makefile",
			appName:      "myapp",
			packageName:  "mypackage",
			languageName: "golang",
			templateName: "Makefile.templ",
		},
	}

	t.Log("Give the need to test Makefile creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				makefile := golang.NewMakeFile(tst.appName, tst.packageName)
				if makefile.ApplicationName != tst.appName {
					t.Fatalf("\t%s\tShould have the correct application name : exp[%s] got[%s]\n", failed, tst.appName, makefile.ApplicationName)
				}
				t.Logf("\t%s\tShould have the correct application name\n", succeeded)

				if makefile.PackageName != tst.packageName {
					t.Fatalf("\t%s\tShould have the correct package name : exp[%s] got[%s]\n", failed, tst.packageName, makefile.PackageName)
				}
				t.Logf("\t%s\tShould have the correct package name\n", succeeded)

				if makefile.VersionFileName != golang.DefaultVersionFileName {
					t.Fatalf("\t%s\tShould have the correct version file name : exp[%s] got[%s]\n", failed, golang.DefaultVersionFileName, makefile.VersionFileName)
				}
				t.Logf("\t%s\tShould have the correct version file name\n", succeeded)
			}
		}
	}

	t.Log("Give the need to test correct language information")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				makefile := golang.NewMakeFile(tst.appName, tst.packageName)
				templateInfo := makefile.TemplateFileInfo()
				if templateInfo.Language != tst.languageName {
					t.Fatalf("\t%s\tShould have the correct language name : exp[%s] got[%s]\n", failed, tst.languageName, templateInfo.Language)
				}
				t.Logf("\t%s\tShould have the correct language name\n", succeeded)

				if templateInfo.Name != tst.templateName {
					t.Fatalf("\t%s\tShould have the correct template name : exp[%s] got[%s]\n", failed, tst.templateName, templateInfo.Name)
				}
				t.Logf("\t%s\tShould have the correct template name\n", succeeded)
			}
		}

	}
}
