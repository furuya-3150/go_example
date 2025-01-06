package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- 1
	}()

	for {
		select {
		case s := <-ch1:
			fmt.Println(s)
		case <-time.After(1 * time.Second): // ch1が受信できないまま1秒で発動
			fmt.Println("time out")
			return
		/*
		// これがあると無限ループする
		default:
			fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		*/
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