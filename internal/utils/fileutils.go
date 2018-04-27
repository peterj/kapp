package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// CreateFolder creates a folder from the path.
func CreateFolder(path string) error {
	actualPath := ExtractFolder(path)
	if err := os.MkdirAll(actualPath, os.ModePerm); err != nil {
		return errors.Wrap(err, fmt.Sprintf("createFolder: %s", actualPath))
	}
	return nil
}

// ExtractFolder removes the filename from the path and
// returns the folder portion only
func ExtractFolder(path string) string {
	actualPath := path

	// Check if there multiple parts in the path
	if strings.Contains(path, "/") {
		filename := filepath.Base(path)
		actualPath = strings.TrimSuffix(path, filename)
		actualPath = strings.TrimSuffix(actualPath, "/")
	}
	return actualPath
}

// FullPath returns the full path to the template
func FullPath(parts ...string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "fullPath")
	}
	// prepend wd to the parts that were passed in
	return path.Join(append([]string{wd}, parts...)...), nil
}
