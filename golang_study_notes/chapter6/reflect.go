package function

import (
	"fmt"
	"reflect"
)

type S struct{}

type TT struct {
	S
}

func reflect_test() {
	var x int = 5
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Int())
	fmt.Println("value:", v)

}

func (S) sVal()   {}
func (*S) sPtr()  {}
func (TT) tVal()  {}
func (*TT) tPtr() {}

func reflect_funciton(a interface{}) {
	t := reflect.TypeOf(TT{})
	fmt.Println(t)
	fmt.Println(t.NumMethod())
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

//func main() {
//	reflect_funciton(TT{})
//
//}
