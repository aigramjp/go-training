package main

import (
	"fmt"
)

func main() {
	//func_pointer()
	//func_new()
	func_struct()
}

type Vertex struct {
	// 大文字にすることで、外部から呼び出すことができる
	X, Z int
	Y    int
}

func func_struct() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)

}

func func_new() {
	// newがないとメモリに確保されないので、注意
	// ポインタが返り値になるようなものにNewを使用し、それ以外にMakeを使用する
	// channelやstructでもポインタが帰ってくる
	var p *int = new(int)
	fmt.Println(p)
}

func func_pointer() {
	var n int = 100
	one(&n)
	fmt.Println(n)
}

func one(x *int) {
	*x = 1
}
