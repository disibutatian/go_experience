package inter

import (
	"fmt"
	"testing"
)

func Test_Interface(t *testing.T) {
	testIPAddr()
	fmt.Printf("------------------\n")
	testIPAddrNew()
	fmt.Printf("------------------\n")
	testInterfaceNil()
	fmt.Printf("------------------\n")
	testIterfaceVariable()
	fmt.Printf("------------------\n")
}
