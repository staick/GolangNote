package sub

import "fmt"

func init() {
	fmt.Println("this is the uitils in the package sub")
}

func test4() {
	Sub(2, 1) //由于Sub和test4.go在同一个包下面，所以可以只用，并且不需要sub.形式
}
