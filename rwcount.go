// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package rwcount implements a simple counter that wraps an io.Reader or io.Writer.
// Userful for functions (like binary.Read/Write) which do not return read/write counts.
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

// Sets given argument to number of bytes read. Pointer argument allows for it to be used
// with defer and named arguments
func (c CountReader) BytesRead(d *int64) {
	*d = c.bytesRead
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

// Sets given argument to number of bytes written. Pointer argument allows for it to be used
// with defer and named arguments
func (c CountWriter) BytesWritten(d *int64) {
	*d = c.bytesWritten
}
