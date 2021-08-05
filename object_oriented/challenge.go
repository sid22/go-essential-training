// print a comment here
package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wtr io.Writer
}

func (c Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A') // 'a' letters have higher ascii value than capital letters

	out := make([]byte, len(p)) // we make a buffer of type byte to put the updated string in, make is used to make a slice it needs however in sync the []type not only type
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.wtr.Write(out)
}

func main() {
	// buff := []byte{'a', 'B', 'd', 'c'}
	cp := &Capper{os.Stdout}
	// cp.Write(buff)
	fmt.Fprintln(cp, "hello woRld")

}
