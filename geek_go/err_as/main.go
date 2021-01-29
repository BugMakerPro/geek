package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

func (ce commonError) Error() string {
	return ce.errorMsg
}

func main() {
	err := commonError{
		errorCode: 0,
		errorMsg:  "err msg",
	}
	newErr := errors.Wrap(err, "wrap commonError")
	cm := commonError{}
	if errors.As(newErr, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	}
}
