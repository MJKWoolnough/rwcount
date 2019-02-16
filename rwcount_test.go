package rwcount

import (
	"testing"
)

type null struct{}

func (null) Read(p []byte) (int, error) {
	return len(p), nil
}

func (null) Write(p []byte) (int, error) {
	return len(p), nil
}

func TestCounts(t *testing.T) {
	data := make([]byte, 0, 1500)
	r := Reader{Reader: null{}}
	w := Writer{Writer: null{}}
	var total int64
	for i := 0; i < 1500; i++ {
		total += int64(i)
		n, _ := w.Write(data[:i])
		if n != i {
			t.Errorf("failed to write %d bytes to buffer, wrote %d", i, n)
		} else if total != w.Count {
			t.Errorf("written bytes returned (%d) didn't match expected (%d)", w.Count, total)
		}
		n, _ = r.Read(data[:i])
		if n != i {
			t.Errorf("failed to read %d bytes from buffer, read %d", i, n)
		} else if total != r.Count {
			t.Errorf("read bytes returned (%d) didn't match expected (%d)", r.Count, total)
		}
	}
}
