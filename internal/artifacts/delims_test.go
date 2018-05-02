package artifacts_test

import (
	"testing"

	. "github.com/peterj/kapp/internal/artifacts"
)

func TestDelims(t *testing.T) {
	tt := []struct {
		testName   string
		delimLeft  string
		delimRight string
	}{
		{
			testName:   "custom delims",
			delimLeft:  "[[",
			delimRight: "]]",
		},
	}

	t.Log("Give the need to test default delims creation")
	{
		defaultDelims := NewDefaultDelims()
		if defaultDelims.Left != "{{" {
			t.Fatalf("\t%s\tShould have the correct left delimiters : exp[%s] got[%s]\n", failed, "{{", defaultDelims.Left)
		}
		t.Logf("\t%s\tShould have the correct left delimiters\n", succeeded)

		if defaultDelims.Right != "}}" {
			t.Fatalf("\t%s\tShould have the correct right delimiters : exp[%s] got[%s]\n", failed, "}}", defaultDelims.Right)
		}
		t.Logf("\t%s\tShould have the correct right delimiters\n", succeeded)
	}

	t.Log("Give the need to test custom delims creation")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				delims := NewDelims(tst.delimLeft, tst.delimRight)
				if delims.Left != tst.delimLeft {
					t.Fatalf("\t%s\tShould have the correct left delimiters : exp[%s] got[%s]\n", failed, tst.delimLeft, delims.Left)
				}
				t.Logf("\t%s\tShould have the correct left delimiters\n", succeeded)

				if delims.Right != tst.delimRight {
					t.Fatalf("\t%s\tShould have the correct right delimiters : exp[%s] got[%s]\n", failed, tst.delimRight, delims.Right)
				}
				t.Logf("\t%s\tShould have the correct right delimiters\n", succeeded)
			}
		}
	}
}
