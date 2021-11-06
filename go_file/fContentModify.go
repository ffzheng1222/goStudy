package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func FindMDfile(fSlice []string) ([]string, error) {
	resultMDSlice := make([]string, 0)
	for _, fName := range fSlice {
		if strings.HasSuffix(fName, "md") {
			resultMDSlice = append(resultMDSlice, fName)
		}
	}
	return resultMDSlice, nil
}


func ModifyFileInfo(mdf []string) {
    for _, mdfNameStr := range mdf {
        ModifyPngPath(mdfNameStr)
        //return
    }
}


func ModifyPngPath(mdfileName string) {
	in, err := os.Open(mdfileName)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	out, err := os.OpenFile(mdfileName+"#bak", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()

	reader := bufio.NewReader(in)
	for {
		lineStr, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.Contains(lineStr, REPLACEPNGPATH) {
			newLine := strings.Replace(lineStr, REPLACEPNGPATH, GITHUBPNGPATH, -1)
			//fmt.Printf("tony old: %v\n", lineStr)
			fmt.Printf("tony new: %v\n", newLine)
			lineStr = newLine
		}
		_, err = out.WriteString(lineStr + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
	}
}

