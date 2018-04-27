package artifacts_test

import (
	"os"
	"path"
	"testing"

	. "github.com/peterj/create-k8s-app/internal/artifacts"
)

const succeeded = "\u2713"
const failed = "\u2717"

func TestProjectFile(t *testing.T) {
	wd, _ := os.Getwd()

	tt := []struct {
		testName                 string
		templateName             string
		targetFileNamePath       string
		language                 string
		expectedTemplateFullPath string
	}{
		{
			testName:                 "create ProjectFile instance",
			templateName:             "Makefile.templ",
			targetFileNamePath:       "somefile.txt",
			language:                 "golang",
			expectedTemplateFullPath: path.Join(wd, "templates", "golang", "Makefile.templ"),
		},
	}

	t.Log("Give the need to test project file creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				projectFile := NewProjectFile(tst.templateName, tst.targetFileNamePath, tst.language)
				if projectFile.TemplateName != tst.templateName {
					t.Fatalf("\t%s\tShould have the correct template name : exp[%s] got[%s]\n", failed, tst.templateName, projectFile.TemplateName)
				}
				t.Logf("\t%s\tShould have the correct template name\n", succeeded)

				if projectFile.TargetFileNamePath != tst.targetFileNamePath {
					t.Fatalf("\t%s\tShould have the correct target filename path : exp[%s] got[%s]\n", failed, tst.targetFileNamePath, projectFile.TargetFileNamePath)
				}
				t.Logf("\t%s\tShould have the correct target filename path\n", succeeded)

				if projectFile.Language != tst.language {
					t.Fatalf("\t%s\tShould have the correct language : exp[%s] got[%s]\n", failed, tst.language, projectFile.Language)
				}
				t.Logf("\t%s\tShould have the correct language\n", succeeded)

				actualTemplateFullPath, _ := projectFile.GetTemplateFullPath()
				if actualTemplateFullPath != tst.expectedTemplateFullPath {
					t.Fatalf("\t%s\tShould have the correct template path : exp[%s] got[%s]\n", failed, tst.expectedTemplateFullPath, actualTemplateFullPath)
				}
				t.Logf("\t%s\tShould have the correct template path\n", succeeded)

			}
		}
	}
}
