package artifacts_test

import (
	"testing"

	. "github.com/peterj/kapp/internal/artifacts"
)

const succeeded = "\u2713"
const failed = "\u2717"

func TestProjectFile(t *testing.T) {
	tt := []struct {
		testName           string
		template           string
		targetFileNamePath string
		delimLeft          string
		delimRight         string
	}{
		{
			testName:           "create ProjectFile instance",
			template:           "Contents of the template are here",
			targetFileNamePath: "somefile.txt",
			delimLeft:          "{{",
			delimRight:         "}}",
		},
	}

	t.Log("Give the need to test project file creation")
	{
		defaultDelims := NewDefaultDelims()
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				projectFile := NewProjectFile(tst.template, tst.targetFileNamePath, defaultDelims)
				if projectFile.Template != tst.template {
					t.Fatalf("\t%s\tShould have the correct template contents : exp[%s] got[%s]\n", failed, tst.template, projectFile.Template)
				}
				t.Logf("\t%s\tShould have the correct template contents\n", succeeded)

				if projectFile.TargetFileNamePath != tst.targetFileNamePath {
					t.Fatalf("\t%s\tShould have the correct target filename path : exp[%s] got[%s]\n", failed, tst.targetFileNamePath, projectFile.TargetFileNamePath)
				}
				t.Logf("\t%s\tShould have the correct target filename path\n", succeeded)

				if projectFile.Delimiters.Left != tst.delimLeft {
					t.Fatalf("\t%s\tShould have the correct left delimiters : exp[%s] got[%s]\n", failed, tst.delimLeft, projectFile.Delimiters.Left)
				}
				t.Logf("\t%s\tShould have the correct left delimiters\n", succeeded)

				if projectFile.Delimiters.Right != tst.delimRight {
					t.Fatalf("\t%s\tShould have the correct right delimiters : exp[%s] got[%s]\n", failed, tst.delimRight, projectFile.Delimiters.Right)
				}
				t.Logf("\t%s\tShould have the correct right delimiters\n", succeeded)

			}
		}
	}
}
