package main

import "fmt"


func noEmpty(color []string) []string {
    var strOut = make([]string, 0)

    for _, str := range color {
        if str != "" {
            strOut = append(strOut, str)
        }
    }

    return strOut
}


func main() {
    var colorArray = [...]string{"red", "blue", "", "yellow", "green", "", ""}

    colorSlice := colorArray[:]
    fmt.Printf("colorSlice: %q\n", colorSlice)
    //fmt.Println("colorSlice:", colorSlice)

    dealColor := noEmpty(colorSlice)
    //fmt.Println("dealColor:", dealColor)
    fmt.Printf("dealColor: %q\n", dealColor)
}
