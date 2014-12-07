// Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
// Useful for functions (like binary.Read/Write) which do not return read/write counts.
package rwcount

import "io"

// Used to wrap a io.Reader for counting
type CountReader struct {
	io.Reader
	bytesRead int64
}

func (c *CountReader) Read(d []byte) (total int, err error) {
	total, err = c.Reader.Read(d)
	c.bytesRead += int64(total)
	return
}

// Returns the number of bytes read
func (c CountReader) BytesRead() int64 {
	return c.bytesRead
}

// Used to wrap a io.Writer for counting
type CountWriter struct {
	io.Writer
	bytesWritten int64
}

func (c *CountWriter) Write(d []byte) (total int, err error) {
	total, err = c.Writer.Write(d)
	c.bytesWritten += int64(total)
	return
}

// Returns the number of bytes written
func (c CountWriter) BytesWritten() int64 {
	return c.bytesWritten
}
