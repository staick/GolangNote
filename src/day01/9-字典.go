package main

import "fmt"

func main() {
	//1.定义一个字典
	//学生id ==> 学生名字 idNames
	var idNames map[int]string	//定义一个map，此时这个map是不能直接赋值的，它是空的

	//2.分配空间，使用make，可以不指定长度，但建议直接指定长度，性能更好
	//使用map之前一定要先对map分配空间
	idNames = make(map[int]string, 10)

	//3.定义时直接分配空间
	//idNames := make(map[int]string, 10)		//最常用的方法

	idNames[0] = "duke"
	idNames[1] = "lily"

	//4.遍历map
	for key, value := range idNames {
		fmt.Println("key:", key, ", value:", value)
	}

	//5.如何确定一个key是否存在map中
	//在map中不存在访问越界的问题，它认为所有的key都是有效的，所以访问一个不存在的key不会崩溃，返回这个类型的零值
	//零值： bool->false    数字->0     字符串->空
	name9 := idNames[9]
	fmt.Println("name9:", name9)

	//无法通过获取value来判断一个key是否存在，因此我们需要一个能够校验key是否存在的方式
	value, ok := idNames[1]	//如果id=1是存在的，那么value就是key=1对应的值，ok返回true，反之返回零值，false
	if ok {
		fmt.Println("id=1这个key是存在的， value为：", value)
	}

	value, ok = idNames[10]
	if ok {
		fmt.Println("id=10这个key是存在的，value为：", value)
	} else {
		fmt.Println("id=10这个key不存在，value为：", value)
	}

	//6.删除map中的元素
	//使用自由函数delete来删除指定的key
	delete(idNames, 1)
	fmt.Println("idNames删除后：", idNames)
	delete(idNames, 10)			//删除不存在的元素不会报错
	fmt.Println("idNames删除后：", idNames)

	//并发处理的时候需要对map进行上锁	//TODO
}
