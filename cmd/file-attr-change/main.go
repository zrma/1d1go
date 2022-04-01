package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	targetPath = "Downloads"
	baseFolder = "BaseFolder"
	prefix     = "prefix"
)

func main() {
	files, err := ioutil.ReadDir(targetPath)
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	for _, file := range files {
		if strings.Contains(file.Name(), baseFolder) {
			logger.Info("target->",
				zap.String("path", file.Name()),
			)
			changeModified(file.Name(), logger)
		}
	}
}

func changeModified(rootFile string, logger *zap.Logger) {
	loc := time.FixedZone("UTC+9", 0)
	now := time.Date(2021, 6, 4, 12, 30, 0, 0, loc)
	err := filepath.Walk(filepath.Join(targetPath, rootFile),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.Contains(path, prefix) {
				return nil
			}
			logger.Info("ChTimes",
				zap.String("file", path),
			)
			return os.Chtimes(path, now, now)
		})
	if err != nil {
		logger.Error("err",
			zap.String("file", rootFile),
			zap.Error(err),
		)
	}
}
