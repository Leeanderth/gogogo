package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

// GetCurrentFilePath 获取当前文件路径
func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func main() {
	// err := os.WriteFile("file3.txt", []byte("这是内容"), os.ModePerm)
	// fmt.Println(err)

	wfile, err := os.OpenFile("w.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()
	// wfile.Write([]byte("你好"))

	file, err := io.ReadAll(wfile)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}
