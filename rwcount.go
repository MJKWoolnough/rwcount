// Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
// Useful for functions (like binary.Read/Write) which do not return read/write counts.
package rwcount

import "io"

// CountReader is used to wrap a io.Reader for counting
type CountReader struct {
	io.Reader
	bytesRead int64
	StickyErr bool
	Err       error
}

// Read implements the io.Reader interface
func (c *CountReader) Read(d []byte) (int, error) {
	if c.StickyErr && c.Err != nil {
		return 0, c.Err
	}
	total, err := c.Reader.Read(d)
	c.bytesRead += int64(total)
	c.Err = err
	return total, err
}

// BytesRead returns the number of bytes read
func (c CountReader) BytesRead() int64 {
	return c.bytesRead
}

// CountWriter is used to wrap a io.Writer for counting
type CountWriter struct {
	io.Writer
	bytesWritten int64
	StickyErr    bool
	Err          error
}

// Write implements the io.Writer interface
func (c *CountWriter) Write(d []byte) (int, error) {
	if c.StickyErr && c.Err != nil {
		return 0, c.Err
	}
	total, err := c.Writer.Write(d)
	c.bytesWritten += int64(total)
	c.Err = err
	return
}

// BytesWritten returns the number of bytes written
func (c CountWriter) BytesWritten() int64 {
	return c.bytesWritten
}
