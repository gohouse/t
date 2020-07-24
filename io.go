package t

import (
	"bufio"
	"bytes"
	"io"
)

func ParseReader(arg interface{}) *bufio.Reader {
	switch arg.(type) {
	case bufio.Reader:
		var tmp = arg.(bufio.Reader)
		return &tmp
	case *bufio.Reader:
		return arg.(*bufio.Reader)
	case io.Reader:
		return bufio.NewReader(arg.(io.Reader))
	}
	return bufio.NewReader(bytes.NewBuffer(ParseBytes(arg)))
}

func ParseWriter(arg interface{}) io.Writer {
	switch arg.(type) {
	case bufio.Writer:
		var tmp = arg.(bufio.Writer)
		return &tmp
	case *bufio.Writer:
		return arg.(*bufio.Writer)
	case io.Writer:
		return arg.(io.Writer)
	}
	return nil
}

func (tc TypeContext) Reader() *bufio.Reader {
	return ParseReader(tc.val)
}

func (tc TypeContext) Writer() io.Writer {
	return ParseWriter(tc.val)
}