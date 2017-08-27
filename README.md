# rwcount
--
    import "github.com/MJKWoolnough/rwcount"

Package rwcount implements a simple counter that wraps an io.Reader or
io.Writer. Useful for functions (like binary.Read/Write) which do not return
read/write counts.

## Usage

#### type CountReader

```go
type CountReader struct {
	io.Reader
	BytesRead int64
	Err       error
}
```

CountReader is used to wrap a io.Reader for counting

#### func (*CountReader) Read

```go
func (c *CountReader) Read(d []byte) (int, error)
```
Read implements the io.Reader interface

#### type CountWriter

```go
type CountWriter struct {
	io.Writer
	BytesWritten int64
	Err          error
}
```

CountWriter is used to wrap a io.Writer for counting

#### func (*CountWriter) Write

```go
func (c *CountWriter) Write(d []byte) (int, error)
```
Write implements the io.Writer interface
