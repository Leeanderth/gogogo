package main

import (
	"fmt"
	"sync"
)

func main() {
	var ma sync.Map

	go func() {
		for {
			ma.Store(1, "张三")
		}
	}()
	go func() {
		for {
			fmt.Println(ma.Load(1))
		}
	}()

	select {}
}
