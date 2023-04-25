package main

import (
	"bufio"
	"crypto/md5"
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
	osOpen       = osOpenFunc
)

func mockWalkFunc(path string, f filepath.WalkFunc) error {
	return afero.Walk(fileSystem, path, f)
}

func osOpenFunc(path string) (afero.File, error) {
	return os.Open(path)
}

func mockOpenFunc(path string) (afero.File, error) {
	return fileSystem.Open(path)
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

	type entry struct {
		path string
		hash string
	}
	files := make(map[int64][]entry)

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

		hash := md5.New()
		f, err := osOpen(path)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		if _, err := io.Copy(hash, f); err != nil {
			return err
		}

		files[info.Size()] = append(files[info.Size()], entry{
			path: filepath.ToSlash(path),
			hash: fmt.Sprintf("%x", hash.Sum(nil)),
		})
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	var sameSizes []struct {
		size  int64
		files []entry
	}
	for size, files := range files {
		if len(files) < 2 {
			continue
		}

		sort.Slice(files, func(i, j int) bool {
			return files[i].path < files[j].path
		})

		sameSizes = append(sameSizes, struct {
			size  int64
			files []entry
		}{size, files})
	}

	if sortingOption == "1" {
		sort.Slice(sameSizes, func(i, j int) bool {
			return sameSizes[i].size > sameSizes[j].size
		})
	} else {
		sort.Slice(sameSizes, func(i, j int) bool {
			return sameSizes[i].size < sameSizes[j].size
		})
	}

	for _, r := range sameSizes {
		_, _ = fmt.Fprintf(writer, "%d bytes\n", r.size)
		for _, file := range r.files {
			_, _ = fmt.Fprintf(writer, "%s\n", file.path)
		}
		_, _ = fmt.Fprintln(writer)
	}

	for {
		_, _ = fmt.Fprintln(writer, "Check for duplicates?")

		scanner.Scan()
		sortingOption = scanner.Text()
		if sortingOption == "yes" {
			break
		}
		if sortingOption == "no" {
			return
		}
		_, _ = fmt.Fprintln(writer, "Wrong option")
		_, _ = fmt.Fprintln(writer)
	}

	totIdx := 0
	for _, r := range sameSizes {
		hashMap := make(map[string][]entry)
		for _, file := range r.files {
			hashMap[file.hash] = append(hashMap[file.hash], file)
		}

		var sameHashKeys []string
		for k, v := range hashMap {
			if len(v) > 1 {
				sameHashKeys = append(sameHashKeys, k)
			}
		}

		if len(sameHashKeys) < 1 {
			continue
		}

		sort.Strings(sameHashKeys)

		_, _ = fmt.Fprintf(writer, "%d bytes\n", r.size)
		for _, k := range sameHashKeys {
			files := hashMap[k]
			sort.Slice(files, func(i, j int) bool {
				return files[i].path < files[j].path
			})
			_, _ = fmt.Fprintf(writer, "Hash: %s\n", k)
			for _, file := range files {
				totIdx++
				_, _ = fmt.Fprintf(writer, "%d. %s\n", totIdx, file.path)
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
