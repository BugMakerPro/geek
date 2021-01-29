package main

import (
	"errors"
	"fmt"
)

type errString string

func (e errString) Error() string {
	return string(e)
}
func NewErr(er string) error {
	return errString(er)
}

var (
	ErrNamedType  = NewErr("EOF")
	ErrStructType = errors.New("EOF")
)

func main() {
	if ErrNamedType == NewErr("EOF") {
		fmt.Println("Named Type Err")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Err")
	}
}
// output:
// Named Type Err
