package artifacts_test

import (
	"testing"

	. "github.com/peterj/create-k8s-app/internal/artifacts"
)

const succeeded = "\u2713"
const failed = "\u2717"

func TestProjectFile(t *testing.T) {
	tt := []struct {
		testName           string
		template           string
		targetFileNamePath string
		language           string
	}{
		{
			testName:           "create ProjectFile instance",
			template:           "Contents of the template are here",
			targetFileNamePath: "somefile.txt",
			language:           "golang",
		},
	}

	t.Log("Give the need to test project file creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				projectFile := NewProjectFile(tst.template, tst.targetFileNamePath, tst.language)
				if projectFile.Template != tst.template {
					t.Fatalf("\t%s\tShould have the correct template contents : exp[%s] got[%s]\n", failed, tst.template, projectFile.Template)
				}
				t.Logf("\t%s\tShould have the correct template contents\n", succeeded)

				if projectFile.TargetFileNamePath != tst.targetFileNamePath {
					t.Fatalf("\t%s\tShould have the correct target filename path : exp[%s] got[%s]\n", failed, tst.targetFileNamePath, projectFile.TargetFileNamePath)
				}
				t.Logf("\t%s\tShould have the correct target filename path\n", succeeded)

				if projectFile.Language != tst.language {
					t.Fatalf("\t%s\tShould have the correct language : exp[%s] got[%s]\n", failed, tst.language, projectFile.Language)
				}
				t.Logf("\t%s\tShould have the correct language\n", succeeded)
			}
		}
	}
}
