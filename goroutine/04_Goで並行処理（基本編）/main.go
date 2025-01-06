package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
    go func() {
			// forが回り切った後に上書きされた2で3回出力される
			fmt.Println(i)
    }()
	}
	time.Sleep(time.Second * 1)
}