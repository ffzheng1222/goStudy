package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)



func readBufio(fileName string) {

    file, err := os.Open(fileName)

    if err != nil {
        fmt.Println("open file err !", err)
    }

    defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}


func readIoutil(fileName string)  {
    content, err := ioutil.ReadFile(fileName)

    if err != nil {
        fmt.Println("open file err !", err)
    }

    fmt.Printf("%v", string(content))
    fmt.Printf("\n文件读取结束...\n")
}


func main() {
    fileName := "/root/tony_data/go_project/go_reflection/set_reflection.go"

    // readBufio(fileName)

    readIoutil(fileName)
}
