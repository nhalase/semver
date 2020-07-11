package cmd

import (
	"bytes"
	"fmt"
)

var actual string

func capturePrintfToActual(format string, a ...interface{}) (n int, err error) {
	var buff bytes.Buffer
	n, err = fmt.Fprintf(&buff, format, a...)
	actual = buff.String()
	return n, err
}
