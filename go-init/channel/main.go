package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道
// 若信道为0则最多只能处理一件事，若多个则可以处理多个
// 如：搬家，若信道为0，则A搬了东西则需要等B接手之后才能进行下一步动作。
// 若信道为4，A搬了东西后，可以把东西放在信道中（桌子上），最多放你的长度值

func goShopping(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 购物开始\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	moneyChan <- money
	wait.Done()
}

type UserInfo struct {
	Name  string
	Money int
}

func main() {
	var wait sync.WaitGroup

	nameList := []UserInfo{
		{Name: "张三", Money: 3},
		{Name: "李四", Money: 4},
		{Name: "王五", Money: 5},
		{Name: "赵六", Money: 6},
	}

	startTime := time.Now()
	wait.Add(len(nameList))
	for _, user := range nameList {
		go goShopping(user.Name, user.Money, &wait)
	}

	go func() {
		wait.Wait()
		defer close(moneyChan)
	}()

	var moneyLsit []int
	var sum int
	// for {
	// 	money, ok := <-moneyChan
	// 	fmt.Println(money)
	// 	sum += money
	// 	if !ok {
	// 		break
	// 	}
	// }

	for money := range moneyChan {
		moneyLsit = append(moneyLsit, money)
		sum += money
	}

	fmt.Printf("购物结束,金额%v，总共花费%d，相差%s", moneyLsit, sum, time.Since(startTime))

}
