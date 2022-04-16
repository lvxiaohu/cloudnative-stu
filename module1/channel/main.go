package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Producer 每 1 秒插入 int 类型的数据，队列满时阻塞
func Producer(c chan int) {
	// Ticker 节拍器
	timer := time.NewTicker(time.Second)
	go func() {
		count := 0
		for _ = range timer.C {
			count++
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(100) // n will be between 0 and 10
			fmt.Printf("生产第 %d 个数据：%d \n", count, n)
			c <- n
		}
	}()
}

// Consumer 每秒打印一个数据，队列为空时阻塞
func Consumer(c chan int) {
	timer := time.NewTicker(time.Second)
	go func() {
		count := 0
		for _ = range timer.C {
			count++
			select {
			case v := <-c:
				fmt.Printf("接收到第 %d 个数据：%d \n", count, v)
			case <-timer.C:
				log.Println("接受超时了！！！")
				return
			default:
				fmt.Printf("没有接受到数据～～～")
			}
		}
	}()
}

func main() {
	// 队列： 长度 10 ，元素类型 int
	ch := make(chan int, 10)

	defer close(ch)

	Producer(ch)
	time.Sleep(5 * time.Second)
	Consumer(ch)
	time.Sleep(20 * time.Second)
	fmt.Println("Exit!")
}
