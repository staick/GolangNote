package main

import "fmt"

//Person类，绑定方法：Eat, Run, Laugh, 成员

type Person struct {
	name	string
	age		int
	gender	string
	score	float64
}

//在类外面绑定方法
//Eat使用指针可以修改类属性的值
func (this *Person)Eat() {
	fmt.Println("Person is eating")
	//类的方法，可以使用自己的成员
	fmt.Println(this.name + " is eating!")
}
//Eat2不适用指针，不可以修改类属性的值
func (this Person)Eat2() {
	fmt.Println("Person is eating")
	//类的方法，可以使用自己的成员
	this.name = "Duke"
}

func main() {
	lily := Person{
		name: "Lily",
		age: 18,
		gender: "女生",
		score: 100,
	}
	fmt.Println("lily:", lily)
	lily.Eat()

	lily.Eat2()
	lily.Eat()
}
