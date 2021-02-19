package main

import (
	"embed"
	"fmt"
)

//go:embed .gitignore
var ignoreStr string

//go:embed .gitignore
var ignoreByte []byte

//go:embed static/hello.txt
//go:embed static/hello2.txt
var f embed.FS

func main() {
	fmt.Println(ignoreStr)
	fmt.Println(ignoreByte, string(ignoreByte))

	////go:embed static/hello.txt
	//var hello []byte
	//fmt.Println(hello)

	data, _ := f.ReadFile("static/hello.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("static/hello2.txt")
	fmt.Println(string(data))
}
