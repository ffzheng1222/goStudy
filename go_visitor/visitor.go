package main

import (
    "fmt"
)

// Info 封装了一些client调用时所需要的基本信息
type Info struct {
    // 指定namespace的时候才会设置这个参数
    NameSpace string
    Name      string
    OtherThings  string
}

func (i *Info) Visit(fn VisitorFunc) error {
    fmt.Printf("tony: fn=%p\n", fn)
    return fn(i, nil)
}


type Visitor interface {
    Visit(VisitorFunc) error
}

type  VisitorFunc func(*Info, error) error


//==========================================================================

// Log Visitor
type LogVisitor struct {
    visitor Visitor
}


func (v LogVisitor)Visit(fn VisitorFunc) error  {
    return v.visitor.Visit(func(info *Info, err error) error {
        fmt.Println("LogVisitor() before call function")
        err = fn(info, err)
        if err == nil {
            fmt.Printf("OtherVisitor() Name=%s\n", info.Name)
        }
        fmt.Println("LogVisitor() after  call function")
        return err
    })
}


// Other Visitor: 这个Visitor主要用来访问 Info 结构中的 OtherThings 成员
type OtherVisitor struct {
    visitor Visitor
}

func (v OtherVisitor)Visit(fn VisitorFunc) error  {
    return v.visitor.Visit(func(info *Info, err error) error {
        fmt.Println("OtherVisitor() before call function")
        err = fn(info, err)
        if err == nil {
            fmt.Printf("OtherVisitor() OtherThings=%s\n", info.OtherThings)
        }
        fmt.Println("OtherVisitor() after  call function")
        return err
    })
}



type NameVisitor struct {
    visitor Visitor
}

// name visitor: 假设这个visitor主要用于访问 Info 结构中的 Name 和 NameSpace 成员
func (v NameVisitor)Visit(fn VisitorFunc) error {
    return v.visitor.Visit(func(info *Info, err error) error {
        fmt.Println("NameVisitor() before call function")
        err = fn(info, err)
        if err == nil {
            fmt.Printf("NameVisitor() NameSpace=%s\n", info.NameSpace)
        }
        fmt.Println("NameVisitor() after  call function")
        return err
    })
}


func main()  {
    info := Info{}

    var v Visitor = &info
    v = LogVisitor{v}
    v = OtherVisitor{v}
    v = NameVisitor{v}


    loadFile := func(info *Info, err error) error {
       info.Name = "tonyfan"
       info.NameSpace = "tce"
       info.OtherThings = "We are running as remote team"
       fmt.Println("main: loadFile ...")
       return nil
    }

    v.Visit(loadFile)
}


