# rwcount
--
    import "vimagination.zapto.org/rwcount"

Package rwcount implements a simple counter that wraps an io.Reader or
io.Writer. Useful for functions (like binary.Read/Write) which do not return
read/write counts.

## Usage

#### type Reader

```go
type Reader struct {
	io.Reader
	Count int64
	Err   error
}
```

Reader is used to wrap a io.Reader for counting

#### func (*Reader) Read

```go
func (c *Reader) Read(d []byte) (int, error)
```
Read implements the io.Reader interface

#### type Writer

```go
type Writer struct {
	io.Writer
	Count int64
	Err   error
}
```

CountWriter is used to wrap a io.Writer for counting

#### func (*Writer) Write

```go
func (c *Writer) Write(d []byte) (int, error)
```
Write implements the io.Writer interface
