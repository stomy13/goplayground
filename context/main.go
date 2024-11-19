package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	do()
// }

// func do() {
// 	ctx := context.Background()
// 	ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
// 	defer cancelFunc()

// 	go routine(ctx, 1)
// 	go routine(ctx, 2)
// 	<-ctx.Done()
// 	time.Sleep(time.Second)
// }

// func routine(ctx context.Context, i int) {

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			return
// 		default:
// 			fmt.Println(i)
// 		}
// 	}
// }

func main() {
	// contextを生成
	ctx := context.Background()

	// 親のcontextを生成し、parentに渡す
	ctxParent, cancel := context.WithCancel(ctx)
	go parent(ctxParent, "Hello-parent")

	// parentのcontextをキャンセル。mainを先に終了させないように1秒待ってから終了
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func parent(ctx context.Context, str string) {
	// parentからcontextを生成し、childに渡す
	childCtx, cancel := context.WithCancel(ctx)
	go child(childCtx, "Hello-child")
	defer cancel()
	// 無限ループ
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), str)
			return
		}
	}
}

func child(ctx context.Context, str string) {
	// 無限ループ
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), str)
			return
		}
	}
}
