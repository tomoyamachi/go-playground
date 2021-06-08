package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		cancel()
		switch ctx.Err() {
		case context.DeadlineExceeded:
			log.Println("exceeded")
		}
	}()
	child(ctx)
}

func child(ctx context.Context) {
	sec := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("ctx.Done in child", ctx.Err())
			return
		case <-time.Tick(time.Second):
			sec += 1
			echo(sec)
			if sec > 5 {
				return
			}
		}
	}
}

func echo(num int) {
	fmt.Println("echo", num)
}
