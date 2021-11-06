package main

import "fmt"

func main() {
	//标签 LABEL1
	//goto LABEL1	===> 下次进入循环时，i不会保存之前的状态，重新从0开始计算，重新来过
	//break LABEL1	===> continue会跳到指定的位置，但是会记录之前的状态，i变成1
	//continue LABEL1	===> 直接跳出指定位置的循环

	//标签名是随意取的
	LABEL1:
		for i:=0;i<5;i++ {
			for j:=0;j<5;j++ {
				if j == 3 {
					break LABEL1
				}
				fmt.Println("i:", i, ",j:", j)
			}
		}
}
