package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	list := []int{1, 2}
	fmt.Println(list[2])
}

func main() {
	read()
	fmt.Println("执行xxx")
}
