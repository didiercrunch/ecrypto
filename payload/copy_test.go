package payload

import (
	"bytes"
	"io"
	"testing"
)

type Buffer struct {
	bytes.Buffer
	io.ReaderFrom // conflicts with and hides bytes.Buffer's ReaderFrom.
	io.WriterTo   // conflicts with and hides bytes.Buffer's WriterTo.
}

// Simple tests, primarily to verify the ReadFrom and WriteTo callouts inside Copy and CopyN.

func TestCopy(t *testing.T) {
	rb := new(Buffer)
	wb1 := new(Buffer)
	wb2 := new(Buffer)
	rb.WriteString("hello, world.")
	CopyToTwoWriters(wb1, wb2, rb)
	if wb1.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
	if wb2.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
}
