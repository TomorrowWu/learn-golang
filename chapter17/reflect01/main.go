package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("--start--")
	fmt.Println(s)
	fmt.Println("--end--")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:值为=%v\n", i, val.Field(i))
		//获取struct标签,注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d:tag为=%v\n", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	val.Method(1).Call(nil)

	//调用结构体的第一个方法Method(0)
	params := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(40)}
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

func main() {
	monster := Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
		Sex:   "",
	}
	TestStruct(monster)
}
