package main

import (
	"fmt"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, h *http.Request) {

	fmt.Println(h.URL.Path, h.UserAgent())

	byteData, err := os.ReadFile("index.html")
	if err != nil {
		w.Write([]byte("文件不存在"))
		return
	}

	w.Write(byteData)

}

func main() {

	http.HandleFunc("/", Index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
