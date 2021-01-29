package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// understand  errors.Cause and errors.Wrap
func main() {
	err := getErr()
	fmt.Println(errors.Cause(err))
	fmt.Println(errors.Wrap(err, "err"))
}

func getErr() error {
	return io.EOF
}

//func CountLines(r io.Reader) (int, error) {
//	sc := bufio.NewScanner(r)
//	lines := 0
//	for sc.Scan() {
//		lines++
//	}
//	return lines, sc.Err()
//}
