package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func FindAllsubFile(pathStr string) ([]string, error) {
	fileSlice := make([]string, 0)
	dirInfo, err := ioutil.ReadDir(pathStr)
	if err != nil {
		fmt.Println("read Dir failed !!!")
		return fileSlice, err
	}
	for _, fi := range dirInfo {
		fullName := pathStr + "/" + fi.Name()
		if fi.IsDir() {
			temp, err := FindAllsubFile(fullName)
			if err != nil {
				fmt.Println("read Dir sub file failed !!!")
				fmt.Printf("fullName=%v, err=%v\n", fullName, err)
				return fileSlice, err
			}
			fileSlice = append(fileSlice, temp...)
		} else {
			fileSlice = append(fileSlice, fullName)
		}
	}
	return fileSlice, err
}

func OpFileRename(fnSlice []string, oRstr, nRstr string) {
	for _, filename := range fnSlice {
		var oldFn = &filename
		newFn := new(string)
		FileRename(oldFn, newFn, oRstr, nRstr)
		if *newFn != "" {
			fmt.Printf("%v rename to %v ^_^\n", *oldFn, *newFn)
		}
	}
}

func FileRename(oldFileName, newFileName *string, oRstr, nRstr string) {
	if strings.Contains(*oldFileName, oRstr) {
		// fmt.Println("tony func fileRename:", *oldFileName)
		*newFileName = strings.Replace(*oldFileName, oRstr, nRstr, -1)
		fmt.Println("tony: ", *newFileName)
		_ = os.Rename(*oldFileName, *newFileName)
	}
}

