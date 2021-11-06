package main

import "fmt"

func main() {

	//1-定义，定义一个具有10个数字的数组
	//c语言定义：int nums[10] = {1, 2, 3, 4}
	//go语言定义：nums := [10]int{1, 2, 3, 4}	  (常用方式)
	// var nums = [10]int{1, 2, 3, 4}
	//var nums [10]int= [10]int{1, 2, 3, 4}

	nums := [10]int{1, 2, 3,4}

	//2-遍历，方式一
	for i:=0;i<len(nums); i++ {
		fmt.Println("i:", i, ", j:", nums[i])
	}

	//方式二：for range ===> python支持
	for key, value := range nums {
		fmt.Println("key:", key, ", value:", value)
	}

	//在go语言中，如果想忽略一个值，可以使用_
	//如果两个都忽略，那么不能使用:=，而应该使用 =
	for _, _ = range nums {
		fmt.Println()
	}

	//不定长数组定义
	//3-使用make进行创建数组
}
