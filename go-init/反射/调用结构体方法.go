package main

import (
	"fmt"
	"reflect"
)

type IUser struct {
	Name1 string `big:"-"`
	Name2 string
}

func (i IUser) Call(text string) {
	fmt.Println(i.Name1, text)
}

func applyFunc(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumMethod(); i++ {
		tf := t.Method(i)
		if tf.Name != "Call" {
			continue
		}

		method := v.Method(i)
		method.Call([]reflect.Value{
			reflect.ValueOf("555"),
		})
	}

	v.NumField()

}

// func main() {
// 	stu1 := IUser{
// 		Name1: "1111",
// 		Name2: "name2",
// 	}
// 	applyFunc(&stu1)

// 	// fmt.Println(stu1)
// }
