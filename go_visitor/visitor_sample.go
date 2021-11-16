package  main

import (
    "encoding/json"
    "fmt"
)

type  TonyVisitor func(person Person)

// 基类：被访问的对象，通过accept方法接受访问
type Person interface {
    accept(visitor TonyVisitor)
}


// 存储学生信息的类型，实现了Person接口
type Student struct {
    Name string
    Age int
    Score float64
}

func (s Student)accept(visitor TonyVisitor)  {
    visitor(s)
}


// 存储教师信息，实现了Person接口
type Teacher struct {
    Name string
    Age int
    Course string
}

func (t Teacher)accept(visitor TonyVisitor) {
    visitor(t)
}


/* 定义两个简单的访问器 */
// 导出json格式数据的访问器
func JsonVisitor(person Person)  {
   bytes, err := json.Marshal(person)

   if err != nil {
       panic(err)
   }
   fmt.Printf("\nJsonVisitor: ")
   fmt.Println(string(bytes))
}


// 导出yaml格式信息的访问器
func YamlVisitor(person Person)  {
    bytes, err := json.Marshal(person)

    if err != nil {
        panic(err)
    }
    fmt.Printf("\nYamlVisitor: ")
    fmt.Println(string(bytes))
}


func main()  {
    s := Student{Age:10, Name:"tony", Score:92.5}
    t := Teacher{Name:"lee", Age:28, Course:"math"}

    persons := []Person{s, t}

    for _, person := range persons {
        person.accept(JsonVisitor)
        person.accept(YamlVisitor)
    }
}
