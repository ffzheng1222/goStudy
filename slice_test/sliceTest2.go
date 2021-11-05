package main

import "fmt"


func noRepeat(color []string) []string {
    var strOut = make([]string, 0)

    for _, str := range color {
        i := 0
        for ; i < len(strOut); i++ {
            if strOut[i] == str {
                break
            }
        }
        if i == len(strOut) {
            strOut = append(strOut, str)
        }
    }

    return strOut
}


func main() {
    var colorArray = [...]string{"red", "blue", "yellow", "red", "yellow", "green", "blue", "yellow"}

    colorSlice := colorArray[:]
    fmt.Printf("colorSlice: %q\n", colorSlice)
    //fmt.Println("colorSlice:", colorSlice)

    dealColor := noRepeat(colorSlice)
    //fmt.Println("dealColor:", dealColor)
    fmt.Printf("dealColor: %q\n", dealColor)
}
