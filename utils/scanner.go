package utils

import (
	"io"
	"strings"
)

func NewStringScanner(str string) *StringScanner {
	var data []string
	res := strings.Split(str, "\n")
	for _, line := range res {
		for _, s := range strings.Split(line, " ") {
			s = strings.TrimSpace(s)
			if s != "" {
				data = append(data, s)
			}
		}
	}
	return &StringScanner{data: data}
}

type StringScanner struct {
	data   []string
	cursor int
}

func (s *StringScanner) Scan() bool {
	return s.cursor < len(s.data)
}

func (s *StringScanner) Text() string {
	defer func() { s.cursor++ }()
	return s.data[s.cursor]
}

func (s *StringScanner) Err() error {
	if s.cursor == len(s.data) {
		return io.EOF
	}
	return nil
}
