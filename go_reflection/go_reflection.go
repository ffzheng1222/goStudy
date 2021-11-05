package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json: "name"`
	Age   int     `json: "monster_age"`
	Score float64 `json: "成绩"`
	Sex   string
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float64, sex string)  {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func (s *Monster) SetMonster (name string, age int, score float64, sex string)  {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func (s Monster) Show() {
	fmt.Printf("\n==============start============\n")
	fmt.Println(s)
	fmt.Printf("==============end============\n")
}

func TestStruct1(a interface{}) {
	// 获取a interface接口的数据类型
	interfaceType := reflect.TypeOf(a)
	interfaceValue := reflect.ValueOf(a)

	derivedType := interfaceValue.Kind()
	if derivedType != reflect.Struct {
		fmt.Println("expect struct T_T")
		return
	}

	// 获取struct的所有字段总数
	structFieldNum := interfaceValue.NumField()
	fmt.Printf("%s 结构体有 %d 个字段\n", interfaceType.Name(), structFieldNum)

	// 遍历struct的所有字段
	for i := 0; i < structFieldNum; i++ {
		fmt.Printf("Field %d:  建=%v,  值=%v\n", i, interfaceType.Field(i), interfaceValue.Field(i))

		//获取到struct标签，需要通过reflect.TypeOf来获取tag标签值
		tagValue := interfaceType.Field(i).Tag.Get("json")

		if tagValue != "" {
			fmt.Printf("Field %d:  tag=%v\n", i, tagValue)
		}
	}

	// 获取struct的所有方法总数
	fmt.Println()
	structFuncNum := interfaceValue.NumMethod()
	fmt.Printf("%s 结构体有 %d 个方法\n", interfaceType.Name(), structFuncNum)

	// 调用struct指定方法
	for j := 0; j < structFuncNum; j++ {
		fmt.Printf("func %d:  建=%v,  值=%v\n", j, interfaceType.Method(j), interfaceValue.Method(j))

		//fmt.Printf("tony: %v\n", interfaceType.Method(j).Name)
		switch interfaceType.Method(j).Name {
		case "GetSum":
			var getSumParams []reflect.Value
			getSumParams = append(getSumParams, reflect.ValueOf(10))
			getSumParams = append(getSumParams, reflect.ValueOf(30))
			res := interfaceValue.MethodByName("GetSum").Call(getSumParams)[0].Int()
			fmt.Println("GetSum: ", res)
			fmt.Println()
		case "Set":
			var setParams []reflect.Value
			setParams = append(setParams, reflect.ValueOf("tonyTan"))
			setParams = append(setParams, reflect.ValueOf(18))
			setParams = append(setParams, reflect.ValueOf(93.0))
			setParams = append(setParams, reflect.ValueOf("女"))

			res := interfaceValue.MethodByName("Set").Call(setParams)[0]
			fmt.Printf("func Set: %p\n", &res)
			fmt.Println("Set: ", res)
			fmt.Printf("%T\n", res)
			fmt.Println()
		case "Show":
			res := interfaceValue.MethodByName("Show").Call(nil)
			fmt.Println("Show: ", res)
			fmt.Println()
		}
	}
}

func TestStruct2(a interface{}) {
	// 获取a interface接口的数据类型
	interfaceType := reflect.TypeOf(a)
	interfaceValue := reflect.ValueOf(a)

	// 获取struct的所有字段总数
	structFieldNum := interfaceType.Elem().NumField()
	fmt.Printf("%s 结构体有 %d 个字段\n", interfaceType.Elem().Name(), structFieldNum)

	// 遍历struct的所有字段
	for i := 0; i < structFieldNum; i++ {
		fmt.Printf("Field %d:  建=%v,  值=%v\n", i, interfaceType.Elem().Field(i), interfaceValue.Elem().Field(i))
		//获取到struct标签，需要通过reflect.TypeOf来获取tag标签值
		tagValue := interfaceType.Elem().Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("Field %d:  tag=%v\n", i, tagValue)
		}
	}


	// 获取struct的所有方法总数
	structFuncNum := interfaceValue.NumMethod()
	fmt.Printf("\n%s结构体有 %d 个方法\n", interfaceType.Elem().Name(), structFuncNum)

	// 调用struct指定方法
	for j := 0; j < structFuncNum; j++ {
		fmt.Printf("func %d:  建=%v,  值=%v\n", j, interfaceType.Method(j), interfaceValue.Method(j))

		//fmt.Printf("tony: %v\n", interfaceType.Method(j).Name)
		switch interfaceType.Method(j).Name {
		case "GetSum":
			var getSumParams []reflect.Value
			getSumParams = append(getSumParams, reflect.ValueOf(10))
			getSumParams = append(getSumParams, reflect.ValueOf(30))
			res := interfaceValue.MethodByName("GetSum").Call(getSumParams)[0].Int()
			fmt.Println("GetSum: ", res)
		case "SetMonster":
			var setParams []reflect.Value
			setParams = append(setParams, reflect.ValueOf("tonyTan"))
			setParams = append(setParams, reflect.ValueOf(18))
			setParams = append(setParams, reflect.ValueOf(93.0))
			setParams = append(setParams, reflect.ValueOf("女"))
			interfaceValue.MethodByName("SetMonster").Call(setParams)
		case "Show":
			interfaceValue.MethodByName("Show").Call(nil)
		}
	}
}


func main() {
	//var monster1 = Monster {
	//	Name: "凡凡",
	//	Age: 25,
	//	Score: 98.5,
	//	Sex: "男",
	//}
	//TestStruct1(monster1)

// ===============================

	var monster2 = Monster {
		Name: "小凡",
		Age: 24,
		Score: 97.5,
		Sex: "男",
	}

	fmt.Printf("main Monster: %v\n", monster2)
	TestStruct2(&monster2)
	fmt.Printf("main Monster: %v\n", monster2)
}
