package timeout

import (
	"context"
	"fmt"
	"log"
	"time"
)

func ContextTimeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Hello")
	case <-ctx.Done():
		log.Fatal(ctx.Err().Error())
	}
}
