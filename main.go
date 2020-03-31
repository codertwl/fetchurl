// WebUrls_Spider project main.go
package main

import (
	"flag"
	"fmt"
	"github.com/codertwl/fetchurl/logic"
	"os"
)

func main() {

	outPath := ""
	deep := 2
	flag.StringVar(&outPath, "o", "./", "output dir")
	flag.IntVar(&deep, "d", 1, "-d 1")
	flag.Parse()

	err := os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err:", err)
		return
	}

	name := fmt.Sprintf("%s/urls.txt", outPath)
	file, err := os.Create(name)
	if err != nil {
		panic("创建文件失败:" + name)
	}
	defer file.Close()

	mapUrl := make(map[string]bool, 1000000)
	logic.DoFetch("http://www.hao123.com/", deep, mapUrl, file)
}
