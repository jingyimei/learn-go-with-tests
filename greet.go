package learn_go_with_tests

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) {
	 fmt.Fprintf(w, "Hello, %s", name)
}
