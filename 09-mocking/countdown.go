package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	_, _ = fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
