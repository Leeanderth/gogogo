package main

import (
	"fmt"
	"os"
)

func init() {

	file, err := os.ReadFile("main.gogo")
	if err != nil {
		// log.Fatalln(err)
		panic(err)
	}
	fmt.Println(file)
}

// func main() {
// 	fmt.Println("main")
// }
