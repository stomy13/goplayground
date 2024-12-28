package iotest

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

func readAll(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Test_readAll(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want string
		err  error
	}{
		{name: "read all", r: strings.NewReader("hello"), want: "hello", err: nil},
		{name: "read all with error", r: iotest.ErrReader(io.ErrShortWrite), want: "", err: io.ErrShortWrite},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readAll(tt.r)
			if err != tt.err {
				t.Errorf("readAll() error = %v, want %v", err, tt.err)
			}
			if got != tt.want {
				t.Errorf("readAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
