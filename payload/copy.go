package payload

import (
	"io"
)

type CopyOutput struct {
	Err     error
	Written int64
}

func copyAndUpdate(dst io.Writer, data []byte, copyOutput *CopyOutput) {
	if copyOutput.Err != nil {
		return
	}
	nw, ew := dst.Write(data)
	if nw > 0 {
		copyOutput.Written += int64(nw)
	}
	if ew != nil {
		copyOutput.Err = ew
		return
	}
	if len(data) != nw {
		copyOutput.Err = io.ErrShortWrite
	}
}

func CopyToTwoWriters(dst1 io.Writer, dst2 io.Writer, src io.Reader) (copyOutput1 *CopyOutput, copyOutput2 *CopyOutput) {
	// TODO: if necessary, implements the speedup methods if src implements *WriterTo*
	//       or if both dst implement ReaderFrom.  See
	//       http://golang.org/src/pkg/io/io.go for example
	copyOutput1, copyOutput2 = new(CopyOutput), new(CopyOutput)
	var nr int
	var er error
	buf := make([]byte, 32*1024)
	for {
		nr, er = src.Read(buf)
		if nr > 0 {
			copyAndUpdate(dst1, buf[0:nr], copyOutput1)
			copyAndUpdate(dst2, buf[0:nr], copyOutput2)
		}
		if copyOutput1.Err != nil && copyOutput2.Err != nil {
			break
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			copyOutput1.Err = er
			copyOutput2.Err = er
			break
		}
	}
	return
}
