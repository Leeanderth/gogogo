package main

// func main() {
// 	// file, err := os.Open("hello.txt")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// defer file.Close()
// 	// for {
// 	// 	var byteData = make([]byte, 12)
// 	// 	_, err := file.Read(byteData)
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(string(byteData))
// 	// }

// 	file, err := os.Open("hello.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// buf := bufio.NewReader(file)
// 	// for {
// 	// 	line, _, err := buf.ReadLine()
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(string(line), err)
// 	// }

// 	buf := bufio.NewScanner(file)
// 	// buf.Split(bufio.ScanWords)
// 	// 如果返回true，表示有，继续往下读
// 	for buf.Scan() {
// 		fmt.Println(buf.Text())
// 	}

// }
