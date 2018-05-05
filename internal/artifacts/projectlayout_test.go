package artifacts_test

import (
	"io/ioutil"
	"os"
	"path"
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

	createTests := []struct {
		testName         string
		projectInfo      *ProjectInfo
		projectFiles     []*ProjectFile
		expectedContents string
	}{
		{
			testName: "one project file",
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
					Template:           "{{ .ApplicationName }} {{ .PackageName }} {{ .VersionFileName }} {{ .DockerRepository }}",
				},
			},
			expectedContents: "app package VERSION.txt dockerrepo",
		},
		{
			testName: "multiple project files",
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
					Template:           "{{ .ApplicationName }} {{ .PackageName }} {{ .VersionFileName }} {{ .DockerRepository }}",
				},
				{
					Delimiters: &Delims{
						Left:  "{{",
						Right: "}}",
					},
					TargetFileNamePath: "secondfile",
					Template:           "{{ .ApplicationName }} {{ .PackageName }} {{ .VersionFileName }} {{ .DockerRepository }}",
				},
			},
			expectedContents: "app package VERSION.txt dockerrepo",
		},
		{
			testName: "non-default delimiters, multiple project files",
			projectInfo: &ProjectInfo{
				ApplicationName:  "app",
				PackageName:      "package",
				VersionFileName:  "VERSION.txt",
				DockerRepository: "dockerrepo",
			},
			projectFiles: []*ProjectFile{
				{
					Delimiters: &Delims{
						Left:  "[[",
						Right: "]]",
					},
					TargetFileNamePath: "firstfile",
					Template:           "[[ .ApplicationName ]] [[ .PackageName ]] [[ .VersionFileName ]] [[ .DockerRepository ]]",
				},
				{
					Delimiters: &Delims{
						Left:  "{{",
						Right: "}}",
					},
					TargetFileNamePath: "secondfile",
					Template:           "{{ .ApplicationName }} {{ .PackageName }} {{ .VersionFileName }} {{ .DockerRepository }}",
				},
			},
			expectedContents: "app package VERSION.txt dockerrepo",
		},
	}

	t.Log("Given the need to test project layout Create func")
	{
		for i, tst := range createTests {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				// Change the working folder to a temporary folder
				workingFolder := os.TempDir()
				os.Chdir(workingFolder)

				outputFolder := "testfolder"
				layout, _ := NewProjectLayout(tst.projectInfo, tst.projectFiles)
				layout.Create(outputFolder)

				// Check that the workingFolder + testFolder + targetfilenamepath was created
				for _, f := range tst.projectFiles {
					rootOutputFolder := path.Join(workingFolder, outputFolder)
					targetFilePath := path.Join(rootOutputFolder, f.TargetFileNamePath)
					if _, err := os.Stat(targetFilePath); err != nil {
						t.Fatalf("\t%s\tFile should exist : exp[%s] got[%s]\n", failed, targetFilePath, err.Error())
					}
					t.Logf("\t%s\tCorrect file should exist %s\n", succeeded, f.TargetFileNamePath)

					// Check the file contents
					actualContentsBytes, _ := ioutil.ReadFile(targetFilePath)
					if tst.expectedContents != string(actualContentsBytes) {
						t.Fatalf("\t%s\tFile contents are correct : exp[%s] got[%s]\n", failed, tst.expectedContents, string(actualContentsBytes))
					}
					t.Logf("\t%s\tCorrect file contents\n", succeeded)
				}
			}
		}

	}

	createErrorTests := []struct {
		testName     string
		projectInfo  *ProjectInfo
		projectFiles []*ProjectFile
	}{
		{
			testName: "invalid field in the template (execute template error)",
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
					Template:           "{{ .Missing }}",
				},
			},
		},
	}
	t.Log("Given the need to test Create func errors")
	{
		for i, tst := range createErrorTests {
			t.Logf("\tTest %d: \t%s", i, tst.testName)
			{
				// Change the working folder to a temporary folder
				workingFolder := os.TempDir()
				os.Chdir(workingFolder)

				outputFolder := "testfolder"
				layout, _ := NewProjectLayout(tst.projectInfo, tst.projectFiles)
				err := layout.Create(outputFolder)
				if err == nil {
					t.Fatalf("\t%s\tError should be returned\n", failed)
				}
				t.Logf("\t%s\tError should be returned\n", succeeded)
			}
		}
	}
}
