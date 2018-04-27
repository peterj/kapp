package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/peterj/create-k8s-app/internal/artifacts/golang"

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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("missing folder name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		outputFolder := args[0]
		// If project name is not provided, we just the output folder name
		if ApplicationName == "" {
			ApplicationName = outputFolder
		}

		blue := color.New(color.FgBlue).SprintfFunc()
		green := color.New(color.FgGreen).SprintfFunc()
		fmt.Printf("Creating %s project %s (package: %s)", blue(Language), blue(ApplicationName), blue(PackageName))

		switch strings.ToLower(Language) {
		case "golang":
			{
				layout, err := golang.NewProjectLayout(&golang.ProjectInfo{
					ApplicationName: ApplicationName,
					PackageName:     PackageName,
					VersionFileName: "VERSION.txt",
				})

				if err != nil {
					return err
				}

				if err := layout.Create(outputFolder); err != nil {
					return err
				}
			}
		default:
			{
				return fmt.Errorf("project language %s it not supported", Language)
			}
		}

		wd, _ := os.Getwd()
		fmt.Printf("\n\nDone. Your project is here:\n  %s", green(path.Join(wd, outputFolder)))
		return nil
	},
}
