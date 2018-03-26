package goroutine

import (
	"fmt"
	"testing"
)

func Test_CHANVARIABLE(t *testing.T) {

	fmt.Printf("\n\n")
	fmt.Printf("-------chan_varible-----------\n")
	chan_varible()
	fmt.Printf("------------------\n\n")

	fmt.Printf("-------chan_direction-----------\n")
	chan_direction()
	fmt.Printf("------------------\n\n")

	fmt.Printf("-------chan_sem-----------\n")
	chan_sem()
	fmt.Printf("------------------\n\n")

	//fmt.Printf("-------chan_mutex-----------\n")
	//chan_mutex()
	//fmt.Printf("------------------\n\n")

	//fmt.Printf("-------chan_time-----------\n")
	//chan_time()
	//fmt.Printf("------------------\n\n")
	//fmt.Printf("\n\n")

}

func BenchmarkChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chan_performance_1()
	}
	fmt.Printf("------------------\n")
}

func BenchmarkChanblock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chan_performance_block()
	}
	fmt.Printf("------------------\n")
}
