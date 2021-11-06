package main

import "fmt"

var ROOTPATH string = "/data/gitProject/goStudy/go_file/test_file"
var REPLACEPNGPATH string = "C:\\Users\\13724\\AppData\\Roaming\\Typora\\typora-user-images\\"
var GITHUBPNGPATH string = "https://github.com/ffzheng1222/k8sStudy/blob/master/png/"

func main() {
	/* 文件获取：获取指定目录下所有的文件 (递归获取)  */
	fnSlice, err := FindAllsubFile(ROOTPATH)
	if err != nil {
		fmt.Println("main FindAllsubFile failed !!!")
		return
	}

	/* 文件重命名：将获取到的所有文件名做重命名操作  */
	OpFileRename(fnSlice, " ", "_")
	showFileName(fnSlice)

	/* 文件内容修改：修改*.md文件内部的图片读取路径  */
	mdSlice, err := FindMDfile(fnSlice)
	if err != nil {
		fmt.Println("FindMDfile failed !!!")
		return
	}
	showFileName(mdSlice)
        ModifyFileInfo(mdSlice)
}

func showFileName(flist []string) {
	for _, f := range flist {
		fmt.Printf("filename: %v\n", f)
	}
}

