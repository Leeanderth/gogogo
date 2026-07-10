package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// // Hello 为指定的人返回问候语“.
func Hello(name string) (string, error) {
	// 如果 name 为空，则返回错误.
	if name == "" {
		return "", errors.New("empty name")
	}
	// 返回在消息中嵌入名称的问候语.
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// 传入names 切片，返回map类型数据，key类型是string，value类型是string
func Hellos(names []string) (map[string]string, error) {
	// 使用make函数创建一个map类型数据，key类型是string，value类型是string
	messages := make(map[string]string)
	// for循环遍历names切片，name是切片中的每个元素
	for _, name := range names {
		// 调用Hello函数，传入name，返回message和error
		message, err := Hello(name)
		// 如果error不为空，则返回nil和error
		if err != nil {
			log.Println(err)
			continue
		}
		// 将name和message添加到messages map中
		messages[name] = message
	}
	// 返回messages map和nil
	return messages, nil
}
