package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//1,获取到type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

}

func main() {
	//var num = 100
	//reflectTest01(num)

	excise()

	str := "tom"
	//fs := reflect.ValueOf(str)
	//fs.SetString("jack")

	//要用指针
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")

	fmt.Printf("%v\n", str)
}

func excise() {
	var v float64 = 1.2

	rVal := reflect.ValueOf(v)

	rTyp := rVal.Type()

	rKind := rVal.Kind()

	rFlo := rVal.Float()

	fmt.Printf("rVal=%v,rTyp=%v,rKind=%v,rFlo=%v  \n", rVal, rTyp, rKind, rFlo)

	i := rVal.Interface()

	f, ok := i.(float64)
	if ok {
		fmt.Println("float=", f)
	}
}
