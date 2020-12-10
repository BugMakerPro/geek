package main

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
)

func main()  {
	err:=getErr()
	errors.Cause(err)
	errors.Wrap(err,"err")
}

func getErr()error  {
	return io.EOF
}
func CountLines(r io.Reader)(int,error)  {
	sc:=bufio.NewScanner(r)
	lines:=0
	for sc.Scan(){
		lines++
	}
	return lines,sc.Err()
}
