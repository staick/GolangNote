package main

import "fmt"

type MyInt int //type相当于c语言中的typedef

//go语言结构体type + struct来处理
type Student struct {
	name   string
	age    int
	gender string
	score  int
}

func main() {
	var i, j MyInt

	i, j = 10, 20

	fmt.Println("i+j=", i+j)

	lily := Student{
		name:   "Lily",
		age:    20,
		gender: "女生",
		score:  80,	//最后一个元素后面必须加上逗号，如果不加逗号则必须与}同一行
	}

	fmt.Println("lily:", lily.name, lily.age, lily.gender, lily.score)

	//结构体没有->操作
	s1 := &lily
	fmt.Println("lily使用指针s1.name打印:", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("lily使用指针(*s1)打印:", (*s1).name, s1.age, s1.gender, s1.score)

	//在定义期间对赋值时，如果每个字段都赋值了，那么字段的名字可以省略不屑
	//如果只对局部变量赋值，那么必须明确指定变量名字
	Duke := Student{
		name: "Duke",
		age: 28,
		//"男生",
		//88,
	}
	fmt.Println("Duke:", Duke)
}
