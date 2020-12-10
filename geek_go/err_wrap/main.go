package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	_, err := ReadConfig()
	if err != nil {
		//fmt.Println(err)
		fmt.Printf("original error: %T,%v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
	//sql.ErrNoRows
	//errf()
}

func ReadConfig()([]byte,error)  {
	home:=os.Getenv("home")
	config,err:=ReadFile(filepath.Join(home,"./setting.xml"))
	return config,errors.WithMessage(err,"could not read config")
}

func ReadFile(path string)([]byte,error){
	f,err:=os.Open(path)
	if err!=nil{
		return nil, errors.Wrap(err,"ReadFile() open filed")
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func errf()  {
	err:=fmt.Errorf("formate a err:%w",io.EOF)
	if errors.Is(err,io.EOF){
		fmt.Println("is eof")
	}
}