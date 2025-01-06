package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
    go func(i int) {
			// iをコピーして上書きされないよう修正
			fmt.Println(i)
    }(i)
	}
	time.Sleep(time.Second * 1)
}