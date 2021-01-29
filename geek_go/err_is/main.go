package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// understand  errors.Is
func main() {
	err := fmt.Errorf("formate a err:%w", io.EOF)
	if errors.Is(err, io.EOF) {
		fmt.Println("is eof")
	}
}
