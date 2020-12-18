package main

import (
	"fmt"
	"time"
)

func main() {
	makeRecover()
	time.Sleep(time.Second * 5)
	fmt.Printf("main end")
}
func makeRecover() {
	defer func() {
		fmt.Println("makeRecover() defer")
		if err:=recover();err!=nil{
			fmt.Println("catch panic")
			fmt.Println(err)
		}
	}()
	go makePanic()
}

func makePanic()  {
	defer func() {
		fmt.Println("makePanic() defer1")
	}()
	defer func() {
		fmt.Println("makePanic() defer2")
	}()
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("makePanic catch panic")
			fmt.Println(err)
		}
	}()
	panic("makePanic")
}