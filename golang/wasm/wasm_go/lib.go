package main

import (
	"fmt"
)

func main() {}

//export Hello
func Hello() {
	fmt.Println("hello world")
	//resp, err := http.Get("http://www.baidu.com")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.Status)
}
