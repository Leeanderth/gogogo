package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	connect, err := net.Dial("tcp", "127.0.0.1:3000")

	if err != nil {
		fmt.Println(err)
		return
	}

	// for {
	// 	var byteData = make([]byte, 1024)
	// 	n, err := connect.Read(byteData)
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	fmt.Println(string(byteData[:n]))
	// }

	go func() {
		for {
			var byteData = make([]byte, 1024)
			if err != nil {
				fmt.Println("------", err)
				break
			}
			con, err := connect.Read(byteData)
			if err == io.EOF {
				break
			}
			fmt.Println(string(byteData[:con]))
		}
	}()

	for {
		var text string
		fmt.Println("请输入你的信息：")
		fmt.Scan(&text)

		if text == "q" {
			connect.Close()
			break
		}
		connect.Write([]byte(text))
	}

}
