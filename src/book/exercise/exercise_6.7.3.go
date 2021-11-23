package main

import "fmt"

func main() {
	html := make(map[string]string)
	html["p"] = "段落"
	html["img"] = "图像"
	html["h1"] = "一级标题"
	html["h2"] = "二级标题"

	fmt.Println(html)
}