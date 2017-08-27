// Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
// Useful for functions (like binary.Read/Write) which do not return read/write counts.
package rwcount

import "io"

// CountReader is used to wrap a io.Reader for counting
type CountReader struct {
	io.Reader
	BytesRead int64
	Err       error
}

// Read implements the io.Reader interface
func (c *CountReader) Read(d []byte) (int, error) {
	if c.Err != nil {
		return 0, c.Err
	}
	total, err := c.Reader.Read(d)
	c.BytesRead += int64(total)
	c.Err = err
	return total, err
}

// CountWriter is used to wrap a io.Writer for counting
type CountWriter struct {
	io.Writer
	BytesWritten int64
	Err          error
}

// Write implements the io.Writer interface
func (c *CountWriter) Write(d []byte) (int, error) {
	if c.Err != nil {
		return 0, c.Err
	}
	total, err := c.Writer.Write(d)
	c.BytesWritten += int64(total)
	c.Err = err
	return total, err
}
