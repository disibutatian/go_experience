package reflect

import (
	"fmt"
	"reflect"
	"strings"
)

type manger struct {
	s string
	i int
}

func test_reflect_type() {

	type Kind uint

	const (
		Invalid Kind = iota
		Bool
		Int
	)

	var kindNames = []string{
		Invalid: "invalid",
		Bool:    "bool",
		Int:     "int",
	}

	fmt.Println(len(kindNames))
	fmt.Println(kindNames)

	var a int
	t := reflect.TypeOf(a)
	fmt.Println(t.Name())
	fmt.Println(t.Kind()) // 获取基类型
	//fmt.Println(t.Elem())

	/**************************************/
	//读取结构体字段值
	var m manger
	mp := reflect.TypeOf(m)
	//mp := reflect.TypeOf(&m)
	//fmt.Printf("%T\n", mp)
	fmt.Printf("%v\n", mp)
	//mp = mp.Elem()
	//fmt.Printf("%T\n", mp)
	//fmt.Printf("%v\n", mp)

	for i := 0; i < mp.NumField(); i++ {
		f := mp.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}
}

func test_reflect_interface() {
	m := manger{
		s: "10",
		i: 100,
	}
	ri := reflect.ValueOf(&m)
	if !ri.CanInterface() {
		fmt.Println("can't interface ")
	}

	v, ok := ri.Interface().(*manger) //转换为接口类型，再对变量进行断言
	if !ok {
		fmt.Println("can't interface ")
	}

	v.s = "nihao"
	v.i = 99
	fmt.Println("value change", v)
}

func (manger) Test_function(x, y int) (string, int, int) {
	//fmt.Println("reflect function : ", x, y)
	return "reflect function :", x, y

}
func test_reflect_fuction() {
	var m = manger{
		s: "test",
		i: 66,
	}
	//rm := reflect.ValueOf(&m)
	rm := reflect.ValueOf(m)
	sear_func := rm.MethodByName("Test_function")
	para := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}
	out := sear_func.Call(para)

	for _, x := range out {
		fmt.Println(x)
	}
}

type X struct{}

func (X) Test_function_slice(s string, a ...interface{}) string {
	fmt.Println("type: ", s, "value :", a)
	//fmt.Println("type: ", s, "value :", a)
	return fmt.Sprintf(s, a...)
}

func test_reflect_fuction_slice() {
	var x X

	rm := reflect.ValueOf(&x)
	sear_func := rm.MethodByName("Test_function_slice")
	para := []reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf([]interface{}{"fasd", 6}),
	}
	out := sear_func.CallSlice(para)

	fmt.Println(out)
	//for _, x := range out {
	//	fmt.Println(x)
	//}
}

//func add(x []reflect.Value) (results []reflect.Value) {
func add(x []reflect.Value) (results []reflect.Value) {
	//var ret reflect.Value
	//ret = reflect.ValueOf("fdsa")
	//results = append(results, ret)
	//return

	if len(x) == 0 {
		fmt.Println("args error")
		return nil
	}

	//xtype := reflect.TypeOf(x).Elem() // 获取数组对象的基类型
	var ret reflect.Value

	//fmt.Println(x[0].Kind())
	//获取参数的基类型
	switch x[0].Kind() {
	case reflect.Int:
		var sum = 0
		for _, value := range x {
			//int64() 这是类型转换，不同类型之间只要能够转化就行
			//v.Int() 反射类型的int 值返回，只可以是int 类型，否则panic 报错
			sum += int(value.Int())
		}
		ret = reflect.ValueOf(sum)
		//fmt.Printf("%T\n", ret)
	case reflect.String:
		ss := make([]string, 0, len(x))
		for _, value := range x {
			//ss = append(ss, value)
			ss = append(ss, value.String())
		}

		//ret = reflect.ValueOf(ss)
		ret = reflect.ValueOf(strings.Join(ss, "")) // 需要将切片变为字符串再进行对应的打印

	}

	results = append(results, ret)
	return
}

func makeFunc(ptr interface{}) {
	fn := reflect.ValueOf(ptr).Elem()
	v := reflect.MakeFunc(fn.Type(), add)
	fn.Set(v)
}

func test_reflect_construct() {
	var addint func(x, y int) int
	var addstring func(x, y string) string

	makeFunc(&addint)
	makeFunc(&addstring)

	fmt.Println(addint(3, 4))
	fmt.Println(addstring("hello", "world"))
}
