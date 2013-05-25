# rwcount
--
    import "github.com/MJKWoolnough/rwcount"

Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
Userful for functions (like binary.Read/Write) which do not return read/write counts.

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
func (c CountReader) BytesRead(d *int64)
```
Sets given argument to number of bytes read. Pointer argument allows for it to
be used with defer and named arguments

#### type CountWriter

```go
type CountWriter struct {
	io.Writer
}
```

Used to wrap a io.Writer for counting

#### func (CountWriter) BytesWritten

```go
func (c CountWriter) BytesWritten(d *int64)
```
Sets given argument to number of bytes written. Pointer argument allows for it
to be used with defer and named arguments

