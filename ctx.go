package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	go func(ctx context.Context) {

	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("ctx call successfully.")
	case <-time.After(time.Millisecond * 900):
		fmt.Println("ctx timeout")
	}

	return
}
