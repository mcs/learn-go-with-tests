package main

import (
	"bytes"
	"fmt"
)

func Countdown(out *bytes.Buffer) {
	_, _ = fmt.Fprint(out, "3")
}
