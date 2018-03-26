package pkg

import (
	"fmt"
	"testing"
)

func Test_Package(t *testing.T) {

	test_private()
	fmt.Printf("------------------\n")
	test_internal()
	fmt.Printf("------------------\n")

}
