package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
	}
	_, _ = fmt.Fprint(out, "Go!")
}
func main() {
	Countdown(os.Stdout)
}
