package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json: "name"`
	Age   int
	IsMan bool
}

func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	// NumField 获取字段数量
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		tTag := tf.Tag.Get("json")
		if tTag == "" {
			tTag = tf.Name
		}

		fmt.Println(tTag, v.Field(i))
	}

}

// func main() {
// 	stu1 := Student{
// 		Name:  "李李",
// 		Age:   18,
// 		IsMan: true,
// 	}
// 	ParseJson(stu1)
// }
