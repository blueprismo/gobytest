package poker

import (
	"io"
	"os"
)

// Tape represents an os.File that will re-write from the start on every Write call.
type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, io.SeekStart)
	return t.File.Write(p)
}