package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// import "rsc.io/quote"

func main() {
	// 设置预定义Logger的属性，包括
	// 日志条目前缀和禁用打印的标志
	//  时间、源文件和行号.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"golang", "", "java"}

	messages, err := greetings.Hellos(names)
	// log.Println(messages)

	if err != nil {
		log.Fatal(err)
	}
	// for _, message := range message {
	// 	fmt.Println(message)
	// }
	fmt.Println(messages)
}
