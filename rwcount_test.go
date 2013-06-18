package rwcount

import (
	"bytes"
	"testing"
)

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

func read(from Reader, data []byte) (total int64, read int) {
	cR := &CountReader{Reader: from}
	defer func () { total = cR.BytesRead() }()
	read, _ = cR.Read(data)
	return
}

func write(to Writer, data []byte) (total int64, written int) {
	cW := &CountWriter{Writer: to}
	defer func () { total = cW.BytesWritten() }()
	written, _ = cW.Write(data)
	return
}

func TestCounts(t *testing.T) {
	var (
		buf   []byte
		total int64
		n     int
	)
	b := new(bytes.Buffer)
	for i := 0; i < 1500; i++ {
		buf = make([]byte, i)
		total, n = write(b, buf)
		if n != i {
			t.Errorf("failed to write %d bytes to buffer, wrote %d", i, n)
		} else if total != int64(i) {
			t.Errorf("written bytes returned (%d) didn't match expected (%d)", total, i)
		}
		total, n = read(b, buf)
		if n != i {
			t.Errorf("failed to read %d bytes from buffer, read %d", i, n)
		} else if total != int64(i) {
			t.Errorf("read bytes returned (%d) didn't match expected (%d)", total, i)
		}
	}
}
