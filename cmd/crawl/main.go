package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

var fileSystem = afero.NewOsFs()
var httpClient = &http.Client{Timeout: 30 * time.Second}

func main() {
	if err := createFolder(); err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	const (
		startIndex = 0
		endIndex   = startIndex + 1
	)

	for i := startIndex; i < endIndex; i++ {
		endpoint := fmt.Sprintf("https://nipponsei.minglong.org/index.php?section=Tracker&page=%d", i)
		if err := getPage(endpoint, logger); err != nil {
			logger.Error("fail",
				zap.Int("index", i),
				zap.Error(err),
			)
		}
	}
}

//goland:noinspection SpellCheckingInspection
const fullPath = "D:/Downloads/Nipponsei"

func createFolder() error {
	if err := fileSystem.Mkdir(fullPath, 0755); !os.IsExist(err) {
		return err
	}
	return nil
}

func getPage(endpoint string, logger *zap.Logger) error {
	res, err := httpClient.Get(endpoint)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}
	defer func() { _ = res.Body.Close() }()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	document.Find(".download ").Each(func(i int, s *goquery.Selection) {
		if attr, ok := s.Attr("href"); ok {
			name, err := downloadPathFromURL(attr)
			if err != nil {
				logger.Error("download path",
					zap.Error(err),
				)
				return
			}

			n, err := rand.Int(rand.Reader, big.NewInt(40))
			if err != nil {
				logger.Error("rand",
					zap.Error(err),
				)
			}
			time.Sleep(time.Millisecond*80 + time.Duration(n.Int64()))

			if err := downloadFile(name, attr); err != nil {
				logger.Error("download",
					zap.Error(err),
				)
			}
		}
	})
	return nil
}

//goland:noinspection HttpUrlsUsage
const downloadURLPrefix = "http://tracker.minglong.org/torrents/"

func downloadPathFromURL(downloadURL string) (string, error) {
	if !strings.HasPrefix(downloadURL, downloadURLPrefix) {
		return "", fmt.Errorf("unexpected download URL: %s", downloadURL)
	}

	name, err := url.QueryUnescape(strings.TrimPrefix(downloadURL, downloadURLPrefix))
	if err != nil {
		return "", err
	}

	return safeDownloadPath(name)
}

func safeDownloadPath(name string) (string, error) {
	name = strings.ReplaceAll(name, "?", "-")
	if name == "" || name == "." || name == ".." {
		return "", fmt.Errorf("unsafe download filename: %q", name)
	}
	if filepath.IsAbs(name) || strings.ContainsAny(name, `/\:`) {
		return "", fmt.Errorf("unsafe download filename: %q", name)
	}

	cleanName := filepath.Clean(name)
	if cleanName != name || strings.HasPrefix(cleanName, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("unsafe download filename: %q", name)
	}

	return filepath.Join(fullPath, cleanName), nil
}

func downloadFile(filePath, url string) (err error) {
	out, err := fileSystem.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	return err
}
