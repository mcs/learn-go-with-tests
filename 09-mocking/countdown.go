package main

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	_, _ = fmt.Fprint(out, "3")
}
