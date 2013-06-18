# rwcount
--
    import "github.com/MJKWoolnough/rwcount"

Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
Useful for functions (like binary.Read/Write) which do not return read/write counts.

## Usage

#### type CountReader

```go
type CountReader struct {
	io.Reader
}
```

Used to wrap a io.Reader for counting

#### func (CountReader) BytesRead

```go
func (c CountReader) BytesRead() int64
```
Returns the number of bytes read

#### type CountWriter

```go
type CountWriter struct {
	io.Writer
}
```

Used to wrap a io.Writer for counting

#### func (CountWriter) BytesWritten

```go
func (c CountWriter) BytesWritten() int64
```
Returns the number of bytes written

