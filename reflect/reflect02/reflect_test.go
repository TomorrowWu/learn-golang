package main

import (
	"fmt"
	"reflect"
	"testing"
)

//TestReflectFunc 适配器函数
func TestReflectFunc(t *testing.T) {
	call1 := func(v1, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	//适配器函数作为统一处理接口
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

type user struct {
	UserId string
	Name   string
}

//TestReflectPtr 反射操作结构体
func TestReflectPtr(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)

	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem() //用于set字段值
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())

	//ValueOf()参数,必须为指针,可寻址的value
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickName")
	t.Log("model", model)

}

func TestReflectStruct(t *testing.T) {
	var (
		model user
		sv    reflect.Value
	)

	model = user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())

	//不可寻址value
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickName")
	t.Log("model", model)

}

func TestStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf", st.Kind().String())      //ptr
	st = st.Elem()                                   //st指向的类型
	t.Log("reflect.TypeOf.Elem", st.Kind().String()) //struct
	elem = reflect.New(st)                           //value类型值,持有指向typ的零值的指针

	t.Log("reflet.New", elem.Kind().String())             //ptr
	t.Log("reflet.New.Elem", elem.Elem().Kind().String()) //struct

	model = elem.Interface().(*user) // model的地址和elem一样
	elem = elem.Elem()               //elem指向的值

	elem.FieldByName("UserId").SetString("12345678")
	elem.FieldByName("Name").SetString("nickName")
	t.Log("model model.Name", model, model.Name)

}

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%s 完成了减法运行,%d - %d = %d \n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func TestCal(t *testing.T) {
	c := Cal{
		Num1: 8,
		Num2: 3,
	}

	rVal := reflect.ValueOf(c)

	numField := rVal.NumField()
	for i := 0; i < numField; i++ {
		fmt.Printf("c.Field[%d]=%d \n", i, rVal.Field(i).Int())
	}

	rVal = reflect.ValueOf(&c)
	rVal = rVal.Elem()
	rVal.FieldByName("Num1").SetInt(10)
	rVal.FieldByName("Num2").SetInt(5)

	numMethod := rVal.NumMethod()
	for i := 0; i < numMethod; i++ {
		rVal.Method(i).Call([]reflect.Value{reflect.ValueOf("tom")})
	}
}
