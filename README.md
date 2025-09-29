# rwcount

[![CI](https://github.com/MJKWoolnough/rwcount/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/rwcount/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/rwcount.svg)](https://pkg.go.dev/vimagination.zapto.org/rwcount)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/rwcount)](https://goreportcard.com/report/vimagination.zapto.org/rwcount)

--
    import "vimagination.zapto.org/rwcount"

Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.

Useful for functions (like binary.Read/Write) which do not return read/write counts.

## Highlights

 - Wrap any `io.Reader` or `io.Writer`.
 - Keeps track of bytes read/written.
 - First error is stored and future reads/writes will all return that error.

## Usage

```go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"vimagination.zapto.org/rwcount"
)

func main() {
	var (
		buf    bytes.Buffer
		result uint16
	)

	writer := &rwcount.Writer{Writer: &buf}

	binary.Write(writer, binary.LittleEndian, uint16(1234))

	reader := &rwcount.Reader{Reader: &buf}
	binary.Read(reader, binary.LittleEndian, &result)

	fmt.Printf("Wrote bytes: %d\n"+
		"Write error: %v\n"+
		"Read bytes: %d\n"+
		"Read error: %v\n"+
		"Read value: %d\n",
		writer.Count,
		writer.Err,
		reader.Count,
		reader.Err,
		result,
	)

	// Output:
	// Wrote bytes: 2
	// Write error: <nil>
	// Read bytes: 2
	// Read error: <nil>
	// Read value: 1234
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/rwcount
