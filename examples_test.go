package rwcount_test

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"vimagination.zapto.org/rwcount"
)

func Example() {
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
