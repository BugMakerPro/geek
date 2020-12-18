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

//type errorString struct {
//	s string
//}
//
//func (e errorString) Error() string {
//	return e.s
//}
//func NewError(text string) error {
//	return errorString{text}
//}
//
//var (
//	ErrType = NewError("EOF")
//)
//
//func main() {
//	if ErrType == NewError("EOF") {
//		fmt.Println("Error:", ErrType)
//	}
//}
