package artifacts_test

import (
	"strings"
	"testing"

	"github.com/peterj/create-k8s-app/internal/artifacts"
)

const succeeded = "\u2713"
const failed = "\u2717"

func TestTemplater(t *testing.T) {
	tt := []struct {
		testName       string
		name           string
		language       string
		fullPathSuffix string
	}{
		{
			testName:       "create template file info",
			name:           "mytemplate",
			language:       "golang",
			fullPathSuffix: "/templates/golang/mytemplate",
		},
	}

	t.Log("Give the need to test TemplateFileInfo creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				templateInfo := artifacts.NewTemplateFileInfo(tst.name, tst.language)
				if templateInfo.Name != tst.name {
					t.Fatalf("\t%s\tShould have the correct name : exp[%s] got[%s]\n", failed, tst.name, templateInfo.Name)
				}
				t.Logf("\t%s\tShould have the correct name\n", succeeded)

				if templateInfo.Language != tst.language {
					t.Fatalf("\t%s\tShould have the correct language : exp[%s] got[%s]\n", failed, tst.language, templateInfo.Language)
				}
				t.Logf("\t%s\tShould have the correct language\n", succeeded)

				fullPath, _ := templateInfo.FullPath()
				if !strings.HasSuffix(fullPath, tst.fullPathSuffix) {
					t.Fatalf("\t%s\tShould have the correct full path : exp[%s] got[%s]\n", failed, tst.fullPathSuffix, fullPath)
				}
				t.Logf("\t%s\tShould have the correct full path\n", succeeded)
			}
		}
	}
}
