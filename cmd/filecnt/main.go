package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const targetPath = "D:\\Downloads\\basedir\\"

	countFilesInDir(targetPath)
	searchDeepestFile(targetPath)
}

func countFilesInDir(targetPath string) {
	kvs := make(map[string]int)
	if err := filepath.WalkDir(targetPath, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		parent := filepath.Dir(path)
		parent = strings.Replace(parent, targetPath, "", 1)
		kvs[parent]++

		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	maxKey := ""
	maxVal := 0
	for k, v := range kvs {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}
	log.Println(maxKey, maxVal)
}

func searchDeepestFile(targetPath string) {
	maxDepth := 0
	maxPath := ""
	if err := filepath.WalkDir(targetPath, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		depth := strings.Count(path, "\\")
		if depth > maxDepth {
			maxDepth = depth
			maxPath = path
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	log.Println(filepath.Base(maxPath))
}
