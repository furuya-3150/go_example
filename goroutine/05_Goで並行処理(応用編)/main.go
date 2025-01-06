package main

func main() {
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