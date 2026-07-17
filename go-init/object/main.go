package main

import "fmt"

type Class struct {
	Name string
}

type Student struct {
	Class
	Name string
}

func (s Student) info() {
	fmt.Printf("学生信息，姓名：%s，班级：%s \n", s.Name, s.Class.Name)
}

func (s *Student) setName(name string) {
	s.Name = name
}

func main() {
	c1 := Class{
		Name: "六班",
	}

	s1 := Student{
		Class: c1,
		Name:  "李",
	}

	s1.info()
	s1.setName("李李")
	s1.info()
}
