package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		if('A' <= b[i] && b[i] <= 'M') || ('a' <= b[i] && b[i] <= 'm' ) {
			b[i] = b[i] + 13
		} else if ('N' <= b[i] && b[i] <= 'Z') || ('n' <= b[i] && b[i] <= 'z') {
			b[i] = b[i] - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
