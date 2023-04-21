package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Directory is not specified")
		return
	}

	targetPath := args[1]

	duplicateFileHandler(targetPath, os.Stdout)
}

var fileSystem = afero.NewMemMapFs()

var (
	filepathWalk = filepath.Walk
)

func walkFunc(path string, f filepath.WalkFunc) error {
	return afero.Walk(fileSystem, path, f)
}

func duplicateFileHandler(targetPath string, writer io.Writer) {
	err := filepathWalk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		_, _ = fmt.Fprintln(writer, filepath.ToSlash(path))
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
