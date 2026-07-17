package main

import "fmt"

type Code int

func (c Code) getMessage() string {
	switch c {
	case 1:
		return "成功"
	case 2:
		return "服务端错误"
	case 3:
		return "网络错误"
	}
	return ""
}

func (c Code) value() (code Code, msg string) {
	return c, c.getMessage()
}

func main() {
	var code Code = 1

	code1, msg := code.value()

	fmt.Printf("%v %s\n", code1, msg)
}
