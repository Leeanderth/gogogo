package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")

	if err != nil {
		fmt.Println(err)
		return
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("tcp service: ", addr.String())

	for {
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println("------", err)
			break
		}

		fmt.Println("RemoteAddr", connect.RemoteAddr())
		// connect.Write([]byte("hello word"))

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
			connect.Write([]byte("!" + string(byteData[:con]) + "!"))
		}

		// time.Sleep(2 * time.Second)
		// connect.Close()
	}
}
