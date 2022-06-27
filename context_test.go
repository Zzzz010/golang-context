package belajar_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextD.Value("d")) // bisa
	fmt.Println(contextD.Value("b")) // bisa, karena satu parent
	fmt.Println(contextD.Value("c")) // tidak bisa, karena beda parent
	fmt.Println(contextA.Value("c")) // tidak bisa, induk tidak punya hak mengambil value anaknya
}

func CreateCounter(ctx context.Context) chan int {
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
			}
		}
	}()
	return tujuan
}

func TestContextCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	tujuan := CreateCounter(ctx)
	for n := range tujuan {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())

}
