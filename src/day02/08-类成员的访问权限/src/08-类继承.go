package src

import "fmt"

//在go语言中，权限都是通过首字母的大小写来控制的
//1.import ==> 如果包名不同，那么只有大写字母开头才是public的
//2.对于类里面的成员，方法==> 只有大写开头的才能在其他包中使用

type Human struct {
	Name	string
	Age		int
	Gender	string
}

func (this *Human)Eat() {
	fmt.Println("this is :", this.Name)
}

//定义一个学生类
type Student1 struct {
	Hum		Human	//包含Human类型的变量
	Score	float64
	School	string
}

//定义一个老师，去继承Human
type Teacher struct {
	Human	//直接写Human类型，没有字段名字，就是继承
	Subject	string
}

