package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/afero"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Directory is not specified")
		return
	}

	targetPath := args[1]

	scanner := bufio.NewScanner(os.Stdin)

	duplicateFileHandler(targetPath, scanner, os.Stdout)
}

var fileSystem = afero.NewMemMapFs()

var (
	filepathWalk = filepath.Walk
)

func walkFunc(path string, f filepath.WalkFunc) error {
	return afero.Walk(fileSystem, path, f)
}

func duplicateFileHandler(targetPath string, scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	_, _ = fmt.Fprintln(writer, "Enter file format:")

	scanner.Scan()
	fileFormat := scanner.Text()

	_, _ = fmt.Fprintln(writer, "Size sorting options:")
	_, _ = fmt.Fprintln(writer, "1. Descending")
	_, _ = fmt.Fprintln(writer, "2. Ascending")
	_, _ = fmt.Fprintln(writer)

	var sortingOption string
	for {
		_, _ = fmt.Fprintln(writer, "Enter a sorting option:")

		scanner.Scan()
		sortingOption = scanner.Text()
		if sortingOption == "1" || sortingOption == "2" {
			break
		}
		_, _ = fmt.Fprintln(writer, "Wrong option")
		_, _ = fmt.Fprintln(writer)
	}

	files := make(map[int64][]string)

	err := filepathWalk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if fileFormat != "" && !strings.HasSuffix(info.Name(), fileFormat) {
			return nil
		}

		files[info.Size()] = append(files[info.Size()], filepath.ToSlash(path))
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	var res []struct {
		size  int64
		files []string
	}
	for size, files := range files {
		if len(files) < 2 {
			continue
		}

		sort.Strings(files)

		res = append(res, struct {
			size  int64
			files []string
		}{size, files})
	}

	if sortingOption == "1" {
		sort.Slice(res, func(i, j int) bool {
			return res[i].size > res[j].size
		})
	} else {
		sort.Slice(res, func(i, j int) bool {
			return res[i].size < res[j].size
		})
	}

	for _, r := range res {
		_, _ = fmt.Fprintf(writer, "%d bytes\n", r.size)
		for _, file := range r.files {
			_, _ = fmt.Fprintf(writer, "%s\n", file)
		}
		_, _ = fmt.Fprintln(writer)
	}
}
