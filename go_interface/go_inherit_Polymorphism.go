package main

import "fmt"

type Monkey struct {
    Name string
}

/* 接口声明 */
func (mk * Monkey)Climbing() {
    fmt.Println(mk.Name, " 生来会爬树...")
}

type BirdAble interface {
    Flying()
}

type FishAble interface {
    Swimming()
}


// StudyMonkey 继承 Monkey
type StudyMonkey struct {
    Monkey
}

// 通过StudyMonkey 实现Monkey struct 所定义的接口
func (smk *StudyMonkey)Flying(){
    fmt.Println(smk.Name, " 通过学习，会飞翔...")
}

func (smk *StudyMonkey)Swimming(){
    fmt.Println(smk.Name, " 通过学习，会游泳...")
}

func main() {
    //创建一个StudyMonkey的实例
    monkey := StudyMonkey {
        Monkey {
            Name: "悟空",
        },
    }

    monkey.Climbing()
    monkey.Flying()
    monkey.Swimming()

    //var mon BirdAble
    //mon = &monkey
    //mon.Flying()
}