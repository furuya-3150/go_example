package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case s := <-ch1:
			fmt.Println(s)
		case <- t.C:
			fmt.Println("time out by ticker")
			return
		default:
			fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func restFunc() <-chan int {
	// 1. チャネルを定義
	result := make(chan int)

	// 2. ゴールーチンを立てて
	go func() {
		defer close(result) // 4. closeするのを忘れずに

		// 3. その中で、resultチャネルに値を送る処理をする
		// (例)
		for i := 0; i < 5; i++ {
			result <- 1
		}
	}()

	// 5. 返り値にresultチャネルを返す
	return result
}