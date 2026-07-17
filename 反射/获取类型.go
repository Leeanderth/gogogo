package main

import (
	"fmt"
	"reflect"
)

func GetType(obj any) {
	// 获取类型
	t := reflect.TypeOf(obj)

	// 获取类型的值（看看是什么类型）
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("数字")
	case reflect.String:
		fmt.Println("字符串")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.Array:
		fmt.Println("数组")
	case reflect.Slice:
		fmt.Println("切片")
	}

}

// func main() {

// 	getType("你好")
// 	getType(123)
// 	getType(struct {
// 		Name string
// 	}{
// 		Name: "1111",
// 	})
// 	getType([2]int{1, 2})
// 	getType([]int{1, 2, 3})
// }
