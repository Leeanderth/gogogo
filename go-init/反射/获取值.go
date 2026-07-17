package main

import (
	"fmt"
	"reflect"
)

func getValue(obj any) {
	// 获取类型
	v := reflect.ValueOf(obj)

	// 获取类型的值（看看是什么类型）
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("数字", v.Int())
	case reflect.String:
		fmt.Println("字符串", v.String())
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.Array:
		fmt.Println("数组")
	case reflect.Slice:
		fmt.Println("切片")
	}

}

// func main() {

// 	// GetType("你好")
// 	// GetType(123)
// 	// GetType(struct {
// 	// 	Name string
// 	// }{
// 	// 	Name: "1111",
// 	// })
// 	// GetType([2]int{1, 2})
// 	// GetType([]int{1, 2, 3})
// 	getValue("你好")
// 	getValue(123)
// 	getValue(struct {
// 		Name string
// 	}{
// 		Name: "1111",
// 	})
// 	getValue([2]int{1, 2})
// 	getValue([]int{1, 2, 3})
// }
