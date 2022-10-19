package utils

func NewStringWriter() *StringWriter {
	return &StringWriter{}
}

type StringWriter struct {
	buffer []byte
	result string
}

func (s *StringWriter) Write(p []byte) (n int, err error) {
	s.buffer = append(s.buffer, p...)
	return len(p), nil
}

func (s *StringWriter) Flush() error {
	s.result = string(s.buffer)
	s.buffer = nil
	return nil
}

func (s *StringWriter) String() string {
	return s.result
}
