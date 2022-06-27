package belajar_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter2(ctx context.Context) chan int {
	tujuan := make(chan int)
	go func() {
		defer close(tujuan)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				tujuan <- counter
				counter++
				time.Sleep(1 * time.Second) // slow down simulation
			}
		}
	}()
	return tujuan
}

func TestContextTimerCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	tujuan := CreateCounter2(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range tujuan {
		fmt.Println("Counter", n)
		// if n == 10 {
		// 	break
		// }

	}

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())

}
