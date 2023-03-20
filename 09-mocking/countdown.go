package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	_, _ = fmt.Fprint(out, "3\n2\n1\nGo!")
}

func main() {
	Countdown(os.Stdout)
}
