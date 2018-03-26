package function

import (
	"fmt"
	"testing"
)

type test_example struct {
	in  *N
	out string
}

func Test_ToString(t *testing.T) {
	// t.Fatal("not implemented")
	var k N = 5
	var n *N = &k
	te := []test_example{
		{n, "val:5"},
		{n, "val:5"},
	}

	for _, value := range te {
		//fmt.Println(value)
		if value.in.toString() != value.out {
			t.Error("error")
		}
	}

	fmt.Printf("------------------\n")
}

func Test_Reflect(t *testing.T) {
	reflect_test()

	//t := T
	//var k T
	k := TT{}
	reflect_funciton(k)
	reflect_funciton(&k)
	fmt.Printf("------------------\n")
}

func Test_Express(t *testing.T) {
	var n N = 8

	f := N.express
	f(n)

	f2 := (*N).express
	n += 1
	f2(&n)
	fmt.Printf("------------------\n")
}

func Test_value(t *testing.T) {
	var n N = 8

	f1 := n.value
	n += 1
	f2 := n.value

	f3 := (&n).value_prt
	n += 1
	f4 := (&n).value_prt

	f1()
	f2()
	f3()
	f4()
	//(*N)(nil).value_prt()

	fmt.Printf("------------------\n")
}
