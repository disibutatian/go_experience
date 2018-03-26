package function

import (
	"fmt"
)

type N int

func (n *N) toString() string {
	return fmt.Sprintf("val:%v", *n)
}

func (n N) express() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func (n N) value() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func (n *N) value_prt() {
	fmt.Printf("test.n: %p, %d\n", n, *n)
}
