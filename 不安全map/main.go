package main

import "fmt"

func main() {
	ma := map[int]string{
		1: "1",
	}

	go func() {
		for {
			ma[2] = "张三"
		}
	}()
	go func() {
		for {
			fmt.Println(ma[2])
		}
	}()

	select {}
}
