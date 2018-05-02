package utils_test

import (
	"os"
	"path"
	"testing"

	. "github.com/peterj/kapp/internal/utils"
)

const succeeded = "\u2713"
const failed = "\u2717"

func TestFileUtils(t *testing.T) {
	tt := []struct {
		testName       string
		path           string
		expectedFolder string
	}{
		{
			testName:       "single folder",
			path:           "foldername",
			expectedFolder: "foldername",
		},
		{
			testName:       "single folder with '/'",
			path:           "foldername/",
			expectedFolder: "foldername",
		},
		{
			testName:       "single folder with a file",
			path:           "foldername/myfile.txt",
			expectedFolder: "foldername",
		},
		{
			testName:       "multiple folders with a file",
			path:           "foldername/subfolder/myfile.txt",
			expectedFolder: "foldername/subfolder",
		},
		{
			testName:       "no folder (empty path)",
			path:           "",
			expectedFolder: "",
		},
	}

	t.Log("Give the need to test extract folder functionality")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				actualFolder := ExtractFolder(tst.path)
				if actualFolder != tst.expectedFolder {
					t.Fatalf("\t%s\tShould have the correct folder name : exp[%s] got[%s]\n", failed, tst.expectedFolder, actualFolder)
				}
				t.Logf("\t%s\tShould have the correct folder name\n", succeeded)
			}
		}
	}

	wd, _ := os.Getwd()
	fullPathTests := []struct {
		testName     string
		parts        []string
		expectedPath string
	}{
		{
			testName:     "multiple folders",
			parts:        []string{"one", "two", "three"},
			expectedPath: path.Join(wd, "one", "two", "three"),
		},
		{
			testName:     "single folder",
			parts:        []string{"one"},
			expectedPath: path.Join(wd, "one"),
		},
		{
			testName:     "empty (no folder)",
			parts:        []string{""},
			expectedPath: path.Join(wd),
		},
	}

	t.Log("Give the need to test full path functionality")
	{
		for i, tst := range fullPathTests {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				fullPath, _ := FullPath(tst.parts...)

				if fullPath != tst.expectedPath {
					t.Fatalf("\t%s\tShould have the correct path : exp[%s] got[%s]\n", failed, tst.expectedPath, fullPath)
				}
				t.Logf("\t%s\tShould have the correct path\n", succeeded)
			}
		}
	}

	createFolderTests := []struct {
		testName       string
		inputPath      string
		expectedFolder string
	}{
		{
			testName:       "single folder, single file",
			inputPath:      "somefolder/file.txt",
			expectedFolder: "somefolder",
		},
		{
			testName:       "single folder",
			inputPath:      "myfolder",
			expectedFolder: "myfolder",
		},
		{
			testName:       "multiple folders",
			inputPath:      "subfold1/subfold2/",
			expectedFolder: "subfold1/subfold2",
		},
		{
			testName:       "multiple folders with a file",
			inputPath:      "subfold1/subfold2/myfile.txt",
			expectedFolder: "subfold1/subfold2",
		},
	}
	t.Log("Given the need to test create folder functionality")
	{
		for i, tst := range createFolderTests {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				tempFolder := os.TempDir()
				os.Chdir(tempFolder)
				CreateFolder(tst.inputPath)

				pathToTest := path.Join(tempFolder, tst.expectedFolder)
				_, err := os.Stat(pathToTest)
				if err != nil {
					t.Fatalf("\t%s\tFolder should exist : exp[%s] got[%s]\n", failed, tst.expectedFolder, pathToTest)
				}
				t.Logf("\t%s\tCorrect folder should exist\n", succeeded)
			}
		}
	}
}
