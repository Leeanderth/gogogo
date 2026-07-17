package main

import (
	"fmt"
	"reflect"
)

func setValue(obj any, value any) {
	// 获取类型
	v1 := reflect.ValueOf(obj)
	v2 := reflect.ValueOf(value)

	// 因传入的是指针，所以需要解引用之后在进行对比类型
	if v1.Elem().Kind() != v2.Kind() {
		fmt.Println("修改失败，参数类型不同")
		return
	}

	// 获取类型的值（看看是什么类型）
	switch v1.Elem().Kind() {
	case reflect.Int:
		v1.Elem().SetInt(v2.Int())
		fmt.Println("数字", v1.Elem().Int())
	case reflect.String:
		v1.Elem().SetString(value.(string))
		fmt.Println("字符串", v1.Elem().String())
	}

}

// func main() {
// 	var name string = "李李"
// 	var age int = 123

// 	setValue(&name, "李")
// 	setValue(&age, 12)
// }
