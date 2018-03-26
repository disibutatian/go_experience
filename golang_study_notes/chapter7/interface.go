package inter

import (
	"fmt"
)

type IPAddr [4]byte
type IPAddrNew [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func (i IPAddrNew) String() string {
	return fmt.Sprintf("%v_%v_%v._%v", i[0], i[1], i[2], i[3])
}

func testIPAddr() {
	//m := make(map[string]IPAddr)
	m := map[string]IPAddr{
		"lookback": {127, 0, 0, 1},
		"dns":      {127, 2, 0, 1},
	}

	for key, ipaddr := range m {
		fmt.Printf("%v : %v\n", key, ipaddr)
	}
}

func testIPAddrNew() {
	//m := make(map[string]IPAddr)
	m := map[string]IPAddrNew{
		"lookback": {127, 0, 0, 1},
		"dns":      {127, 2, 0, 1},
	}

	for key, ipaddr := range m {
		fmt.Printf("%v : %v\n", key, ipaddr)
	}
}

func testInterfaceNil() {
	var t interface{}
	if t == nil {
		fmt.Printf("empty interface variable is nil\n")
	}

	var t2 interface {
		string() string
	}
	if t2 == nil {
		fmt.Printf("interface variable default is nil\n")
	}
}

type dat struct {
	x int
}

func testIterfaceVariable() {
	var n interface{} = dat{5}

	fmt.Println(n)
	fmt.Println(n.(dat).x)

	var n2 interface{} = &dat{5}

	//fmt.Println(n2.(type))
	fmt.Println(n2)
	fmt.Println(n2.(*dat).x)
}
