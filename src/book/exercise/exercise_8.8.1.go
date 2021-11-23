package main

import "fmt"

type Taxi struct {
	isEmpty bool
	num     int
	isStop  bool
}

type TaxiInterface interface {
	Earn() float64
}

func (taxi *Taxi) Earn() float64 {
	return float64(taxi.num) * 2.5
}

func EarnMoney(t TaxiInterface) float64 {
	return t.Earn()
}

func main() {
	taxi := Taxi{
		isEmpty: false,
		num:     4,
		isStop:  false,
	}

	fmt.Printf("有%d名乘客，挣了%v元钱。", taxi.num, EarnMoney(&taxi))
}