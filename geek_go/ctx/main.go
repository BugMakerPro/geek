package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx,cancel:=context.WithCancel(context.Background())
	go func(c context.Context) {
		time.Sleep(time.Second)
		cancel()
		fmt.Println("canceled")
	}(ctx)
	go func(c context.Context) {
		time.Sleep(time.Second*2)
		select {
		case <-c.Done():
			fmt.Println("func2 done")
		default:
			fmt.Println("func2")
		}
	}(ctx)
	go func(c context.Context) {
		time.Sleep(time.Second*3)
		select {
		case <-c.Done():
			fmt.Println("func3 done")
		default:
			fmt.Println("func3")
		}
	}(ctx)
	time.Sleep(time.Second*5)
}
