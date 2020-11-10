package main

import (
	"fmt"
	// template 是文件夹的名称，因此 go mod里的module名称也要为template，否则会报找
	// 找不到错误，即package template/handlers is not in GOROOT
	controllers "template/handlers"

	jsoniter "github.com/json-iterator/go"
)

// AppInfo example
type AppInfo struct {
	Name string
}

func main() {
	info := AppInfo{
		Name: "GoApp",
	}
	jsonString, _ := jsoniter.Marshal(&info)

	fmt.Println(string(jsonString))
	controllers.ShowInfo()
}
