package test

import (
	"fmt"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	test_test()
	fmt.Println("******************")
}

func TestB(t *testing.T) {
	//expectStr := "test ok"
	expectStr := "test ok1"
	ret := test_test()
	if expectStr != ret {
		t.Errorf("expect %s, got %s", expectStr, ret)
	}
}

func ExampleTest() {
	fmt.Println(test_test())
	//Output:
	//test ok
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
