package main

import "fmt"


func colorCategory(color []string) {
    var strOut = make(map[string]int, 0)

    // 初始化map: strOut
    for _, str := range color {
        strOut[str] = 0
    }

    for _, str := range color {
        for colorKey, colorValue := range strOut {
            if str == colorKey {
                colorValue ++
                strOut[colorKey] = colorValue
            }
        }
    }
    fmt.Println("colorCategory:", strOut)
}


func main() {
    var colorArray = [...]string{"red", "blue", "yellow", "red", "yellow", "green", "blue", "yellow"}

    colorSlice := colorArray[:]
    fmt.Printf("colorSlice: %q\n", colorSlice)
    //fmt.Println("colorSlice:", colorSlice)

    colorCategory(colorSlice)
}
