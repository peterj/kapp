package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/peterj/kapp/internal/artifacts/helm"
	"github.com/pkg/errors"

	"github.com/peterj/kapp/internal/artifacts"

	"github.com/fatih/color"
	"github.com/peterj/kapp/internal/artifacts/dockermake"
	"github.com/peterj/kapp/internal/artifacts/golang"

	"github.com/spf13/cobra"
)

// ApplicationName variable
var ApplicationName string

// PackageName of the project (e.g. github.com/[username]/[projectname])
var PackageName string

// Language represents the project language
var Language string

func init() {
	createCmd.Flags().StringVarP(&ApplicationName, "name", "n", "", "application name")
	createCmd.Flags().StringVarP(&PackageName, "package", "p", "", "package name")
	createCmd.Flags().StringVarP(&Language, "language", "l", "golang", "project language")

	createCmd.MarkFlagRequired("package")
	rootCmd.AddCommand(createCmd)
}

// askYesNoWarning prints a warning yes/no question and waits for user input
func askYesNoWarning(message string) bool {
	red := color.New(color.FgRed).SprintfFunc()
	fmt.Println(red("WARNING!"))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	answer, _ := reader.ReadString('\n')
	return strings.ToLower(string(answer[0])) == "y"
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new Kubernetes app",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("missing folder name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		blue := color.New(color.FgBlue).SprintfFunc()
		green := color.New(color.FgGreen).SprintfFunc()

		outputFolder := args[0]
		wd, _ := os.Getwd()
		appFolder := path.Join(wd, outputFolder)

		if _, err := os.Stat(appFolder); err == nil {
			msg := fmt.Sprintf("Output folder (%s) already exists.\nAre you sure you want to continue (y/n)? ", blue(appFolder))
			if !askYesNoWarning(msg) {
				return fmt.Errorf("Aborted. Output folder already exists")
			}
		}

		// If project name is not provided, we just the output folder name
		if ApplicationName == "" {
			ApplicationName = outputFolder
		}

		if !strings.HasSuffix(PackageName, ApplicationName) {
			msg := fmt.Sprintf("Package name (%s) should contain the application name (%s).\nAre you sure you want to continue (y/n)? ", blue(PackageName), blue(ApplicationName))
			if !askYesNoWarning(msg) {
				return fmt.Errorf("Aborted. Package name did not contain the application name")
			}
		}

		fmt.Printf(`Creating Kubernetes App:
  Application name: %s
  Language........: %s
  Package name....: %s
`, blue(ApplicationName), blue(Language), blue(PackageName))

		switch strings.ToLower(Language) {
		case "golang":
			{
				projectInfo := &artifacts.ProjectInfo{
					ApplicationName: ApplicationName,
					PackageName:     PackageName,
					VersionFileName: "VERSION.txt",
				}

				allFiles := []*artifacts.ProjectFile{}
				for _, e := range golang.ProjectFiles() {
					allFiles = append(allFiles, e)
				}

				for _, e := range helm.ProjectFiles(ApplicationName) {
					allFiles = append(allFiles, e)
				}

				for _, e := range dockermake.ProjectFiles() {
					allFiles = append(allFiles, e)
				}

				layout, err := artifacts.NewProjectLayout(projectInfo, allFiles)
				if err != nil {
					return errors.Wrap(err, "RunE: new project layout")
				}

				if err := layout.Create(outputFolder); err != nil {
					return errors.Wrap(err, fmt.Sprintf("RunE: create layout: output folder '%s'", outputFolder))
				}
			}
		default:
			{
				return errors.Wrap(fmt.Errorf("RunE: project language %s it not supported", Language), "")
			}
		}
		fmt.Printf("Done.")
		fmt.Printf("\n\nGet started:\n  %s\n  %s\n  %s\n", green(fmt.Sprintf("cd %s", appFolder)), green("git init && git add * && git commit -m 'inital commit'"), green("make all"))
		fmt.Printf("\nNote: Refer to README.md (%s) for explanation of how to build and push Docker image and deploy to Kubernetes\n\n", green("https://github.com/peterj/kapp/blob/master/README.md"))

		return nil
	},
}
