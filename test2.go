package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 在 main 函数里调用给定的 rpc 方法，并设置超时时间为 10 秒
// 在等待过程中如果超时则取消等待并打印 "timeout" ，如果没有超时则打印出 rpc 的返回结果。
// rpc 方法不可以修改
func main() {
	ch := make(chan int, 1)
	go func() {
		for {
			x := rpc()
			ch <- x
		}

	}()
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-time.After(time.Duration(10) * time.Second):
			fmt.Println("timeout")
		}
	}
	chBack := make(chan struct{})
	<-chBack

}

// 这是你要调用的方法，可以看作一个黑盒
// 它的耗时是 1~15 秒内的随机数
// 最终返回一个随机的 int 类型
func rpc() int {
	cost := rand.Intn(15) + 1
	fmt.Printf("rpc will cost %d seconds\n", cost)
	time.Sleep(time.Duration(cost) * time.Second)
	return cost
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
