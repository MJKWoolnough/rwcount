// Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
// Useful for functions (like binary.Read/Write) which do not return read/write counts.
package rwcount // import "vimagination.zapto.org/rwcount"

import "io"

// Reader is used to wrap a io.Reader for counting.
type Reader struct {
	io.Reader
	Count int64
	Err   error
}

// Read implements the io.Reader interface.
func (c *Reader) Read(d []byte) (int, error) {
	if c.Err != nil {
		return 0, c.Err
	}

	total, err := c.Reader.Read(d)
	c.Count += int64(total)
	c.Err = err

	return total, err
}

// Writer is used to wrap a io.Writer for counting.
type Writer struct {
	io.Writer
	Count int64
	Err   error
}

// Write implements the io.Writer interface.
func (c *Writer) Write(d []byte) (int, error) {
	if c.Err != nil {
		return 0, c.Err
	}

	total, err := c.Writer.Write(d)
	c.Count += int64(total)
	c.Err = err

	return total, err
}

// WriteString implements the io.StringWriter interface.
func (c *Writer) WriteString(s string) (int, error) {
	if c.Err != nil {
		return 0, c.Err
	}

	total, err := io.WriteString(c.Writer, s)
	c.Count += int64(total)
	c.Err = err

	return total, err
}
