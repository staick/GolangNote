package main

import "fmt"

//在go语言中没有枚举类型，但是可以使用const + iota（常量累加器）来进行模拟
//模拟一个一周的枚举
const (
	MONDAY = iota	//0
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)	//const属于预编译期赋值，不需要:=进行推导

//1.iota是常量组计数器
//2.iota从0开始，每行递增1
//3.常量组有个特点：如果默认不赋值，默认与上一行表达式相同
//4.如果同一行出现两个iota，那么两个iota的值是相同的
//5.每个常量组的iota是独立的，如果遇到const iota会重新清零

func main() {

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)


	//var number int
	//var name string
	//var flag bool
	//
	////可以使用变量组来统一定义变量
	//var (
	//	number int
	//	name string
	//	flag bool
	//)
}
