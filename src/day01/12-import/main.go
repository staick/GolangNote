package main
//不允许出现多个包名

import (
	"day01/12-import/add"
	//SUB "12-import/sub"	//SUB是我们自己重命名的包名
	//. "12-import/sub"		//.代表用户在调用这个包里面的函数时，不需要使用包名.的形式，可以直接使用
	"day01/12-import/sub" //sub是文件夹名，同时也是包名
	"fmt"
)

func main() {
	res := sub.Sub(20, 10)
	fmt.Println("sub(20, 10)", res)

	//add.add(10, 20)无法被调用是应为首字母是小写的，如果一个包里面的函数想对外提供访问权限，那么一定要首字母大写
	//大写字母开头相当于public
	//小写字母开头的函数相当于private，只有相同包名才能使用
	add.Add(1,2)
}
