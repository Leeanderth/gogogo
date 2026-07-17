package main

import (
	"reflect"
	"strings"
)

type User struct {
	Name1 string `big:"-"`
	Name2 string
}

func setStruct(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		tf := t.Field(i)
		tTag := tf.Tag.Get("big")

		if tTag == "" {
			continue
		}
		value.SetString(strings.ToUpper(value.String()))
	}

}

// func main() {
// 	stu1 := User{
// 		Name1: "name1",
// 		Name2: "name2",
// 	}
// 	setStruct(&stu1)

// 	fmt.Println(stu1)
// }
