package reflect

import (
	"fmt"
	"testing"
)

func Test_Package(t *testing.T) {

	test_reflect_type()
	fmt.Printf("------------------\n")

	test_reflect_interface()
	fmt.Printf("------------------\n")
	test_reflect_fuction()
	fmt.Printf("------------------\n")
	test_reflect_fuction_slice()
	fmt.Printf("------------------\n")
	test_reflect_construct()
	fmt.Printf("------------------\n")

}
