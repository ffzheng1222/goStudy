package main

import (
    "fmt"
    "reflect"
)
type User struct{
    Id int
    Name string
    Age int
}
//ToString方法
func (u User) String() string {
    return "User[ Id " + string(u.Id) +"]"
}
//设置Name方法
func (u *User) SetName(name string) string{
    oldName := u.Name
    u.Name = name
    return oldName
}
//年龄数+1
func (u *User) AddAge() bool {
    u.Age++
    return true
}
//测试方法
func (u User) TestUser() {
    fmt.Println("我只是输出某些内容而已....")
}


func TestReflect(a interface{}) {
    var setNameStr string = "SetName"
    var addAgeStr string = "AddAge"

    //1.获取到结构体类型变量的反射类型
    refUser:= reflect.ValueOf(a)  //需要传入指针，后面再解析
    fmt.Println(refUser)

    fmt.Println(refUser.NumMethod())
    fmt.Printf("%s 结构体有 %d 个方法\n", reflect.TypeOf(a).Name(), refUser.NumMethod())
    for j := 0; j < refUser.NumMethod(); j++ {
        fmt.Printf("func %d:  建=%v,  值=%v\n", j, reflect.TypeOf(a).Method(j), refUser.Method(j))
    }

    //2.获取确切的方法名
    //带参数调用方式
    setNameMethod := refUser.MethodByName( setNameStr  )
    args := []reflect.Value{ reflect.ValueOf("Mike")  } //构造一个类型为reflect.Value的切片
    setNameMethod.Call(args) //返回Value类型
    //不带参数调用方式
    addAgeMethod := refUser.MethodByName( addAgeStr )
    addAgeMethod.Call( make([]reflect.Value , 0) )

    //fmt.Println("User.Name = ",*a.Name)
    //fmt.Println("User.Age = ",*a.Age)
}

func main(){
    //通过反射的方式调用结构体类型的方法
    user := User{
        Id : 1,
        Name : "env107" ,
        Age : 18 ,
    }

    TestReflect(&user)
}