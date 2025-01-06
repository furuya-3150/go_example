package main

import (
	"fmt"
	"time"
)

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	// 単一のデータに対して同時に読み書きを行うことで、データの一貫が取れなくなる
	// この例ではdst
	for _, s := range src {
		go func(s int) {

			result := s * 2

			dst = append(dst, result)
		}(s)
	}

	time.Sleep(time.Second)
	fmt.Println(dst)
}