package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

var AppFs = afero.NewOsFs()

func main() {
	if err := createFolder(); err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	const (
		startIdx = 0
		endIdx   = startIdx + 1
	)

	for i := startIdx; i < endIdx; i++ {
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
	if err := AppFs.Mkdir(fullPath, 0755); !os.IsExist(err) {
		return err
	}
	return nil
}

func getPage(endpoint string, logger *zap.Logger) error {
	res, err := http.Get(endpoint)
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
			//goland:noinspection HttpUrlsUsage
			const prefix = "http://tracker.minglong.org/torrents/"
			name, err := url.QueryUnescape(strings.TrimPrefix(attr, prefix))
			if err != nil {
				logger.Error("unescape",
					zap.Error(err),
				)
			}
			name = strings.Replace(name, "?", "_", -1)
			name = fullPath + "/" + name

			time.Sleep(time.Millisecond*80 + time.Duration(rand.Intn(40)))

			if err := downloadFile(name, attr); err != nil {
				logger.Error("download",
					zap.Error(err),
				)
			}
		}
	})
	return nil
}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	resp, err := http.Get(url)
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
