package main

import (
	"errors"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能等于0")
		return
	}

	res = a / b
	return
}

func service() (res int, err error) {
	res, err = div(2, 2)
	if err != nil {
		return
	}

	return
}

// func main() {
// 	res, err := service()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(res)
// }
