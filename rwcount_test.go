package rwcount

import (
	"bytes"
	"testing"
)

func TestCounts(t *testing.T) {
	var (
		buf   []byte
		total int64
		n     int
	)
	b := new(bytes.Buffer)
	for i := 0; i < 1500; i++ {
		cR := &CountReader{Reader: b}
		cW := &CountWriter{Writer: b}
		buf = make([]byte, i)
		n, _ = cW.Write(buf)
		if n != i {
			t.Errorf("failed to write %d bytes to buffer, wrote %d", i, n)
		} else if cW.BytesWritten(&total); total != int64(i) {
			t.Errorf("written bytes returned (%d) didn't match expected (%d)", total, i)
		}
		n, _ = cR.Read(buf)
		if n != i {
			t.Errorf("failed to read %d bytes from buffer, read %d", i, n)
		} else if cR.BytesRead(&total); total != int64(i) {
			t.Errorf("read bytes returned (%d) didn't match expected (%d)", total, i)
		}
	}
}
