package main

import (
	"fmt"
	"time"
)

func main() {
	// モノトニック時間を含むtime.Time値の例
	t := time.Now() // モノトニック時間を含むtime.Time値が返される

	// モノトニック時間を含むことを確認
	fmt.Printf("Time with monotonic clock: %v\n", t)
	fmt.Printf("Unix timestamp: %d\n", t.Unix())
	fmt.Printf("Monotonic clock reading: %d\n", t.UnixNano())

	// 比較演算の例（モノトニック時間が使用される）
	t1 := time.Now()
	time.Sleep(100 * time.Millisecond)
	t2 := time.Now()

	// t2 は必ず t1 より後になる（壁時計が巻き戻されても）
	fmt.Printf("t2 after t1: %v\n", t2.After(t1))
}
