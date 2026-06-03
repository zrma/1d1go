package main

import (
	"path/filepath"
	"testing"
)

func TestDownloadPathFromURLAcceptsSafeTorrentName(t *testing.T) {
	got, err := downloadPathFromURL(downloadURLPrefix + "safe%20name%3F.torrent")
	if err != nil {
		t.Fatalf("downloadPathFromURL returned error: %v", err)
	}

	want := filepath.Join(fullPath, "safe name-.torrent")
	if got != want {
		t.Fatalf("downloadPathFromURL() = %q, want %q", got, want)
	}
}

func TestDownloadPathFromURLRejectsTraversal(t *testing.T) {
	malicious := downloadURLPrefix + "..%2F..%2Foutside.torrent"

	if _, err := downloadPathFromURL(malicious); err == nil {
		t.Fatal("downloadPathFromURL accepted a traversal filename")
	}
}

func TestDownloadPathFromURLRejectsUnexpectedHost(t *testing.T) {
	if _, err := downloadPathFromURL("https://example.com/safe.torrent"); err == nil {
		t.Fatal("downloadPathFromURL accepted an unexpected URL")
	}
}
