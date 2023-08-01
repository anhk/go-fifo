package fifo

import (
	"fmt"
	"testing"
	"time"
)

func TestGoFifo(t *testing.T) {
	f := NewFIFO[string](8)

	for i := 10; i < 1000; i++ {
		f.Push(fmt.Sprintf("%v", i))
		//fmt.Println(f.queue)

		f.Each(func(n string) {
			fmt.Printf("%v ", n)
		})
		fmt.Println()
		time.Sleep(time.Second)
	}
}
