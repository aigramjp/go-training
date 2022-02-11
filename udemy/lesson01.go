package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func init() {
	fmt.Println("init!")
}

func bazz() {
	fmt.Println("Buzz")
}

// 関数の外でも宣言できる(var,変数)
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

// constは型が定義されていない。利用するときに定義される。Overflowするものも保存することができる？その分利用するのにコストが掛かったりする？
const (
	Pi       = 3.14
	Nu       = 1.4
	Username = "test user"
	Password = "test password"
)

func main() {
	//	bazz()
	fmt.Println(time.Now())
	fmt.Println(user.Current())
	fmt.Println(i, f64, s, t, f)

	// 関数の中でしか宣言できない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false

	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Println(Pi, Nu, Username, Password)

	// シフトでははみ出した部分が削除され、反対側から0が格納される
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
	fmt.Println(1 >> 2)

	// Goの文字列は文字配列ではなく、バイト型配列。
	s1 := "Hello world"
	s2 := strings.Replace(s1, "H", "X", 1)
	fmt.Println(s1, ">>", s2)
	fmt.Println(strings.Contains(s1, "world"))
	fmt.Println(s1[0], string(s1[0]))

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println(!false, !true)

	// 文字列から整数に変換
	var s string = "14"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(s, "err")
	} else {
		fmt.Println(s, i+1)
	}

	var a [2]int
	var b [2]int = [2]int{100, 200}
	var c []int = []int{100, 200}
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}

	a[0] = 100
	a[1] = 200
	c = append(c, 300, 400)

	fmt.Println(a, b, c)
	fmt.Println(board)

	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	m := map[string]int{"apple": 100, "banana": 200}
	m["pine"] = 300
	v, ok := m["nothing"]
	fmt.Println(m)
	fmt.Println(v, ok)

	m2 := make(map[string]int)
	m2["pine"] = 300
	fmt.Println(m2)

	r1, r2, r := add(2, 3)
	fmt.Println(r1, r2, r)

	r3 := cal(100, 5)
	fmt.Println(r3)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	c2 := circleArea(3)
	fmt.Println(c1(10), c2(10))

	foo(1, 2, 3)
	foo(1, 2, 3, 4, 5)
	ss := []int{1, 2, 3}
	foo(ss...)
}

func add(x int, y int) (int, int, int) {
	return x, y, x + y
}

func cal(price int, item int) (result int) {
	result = price * item
	return
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func foo(params ...int) {
	fmt.Println(len(params), params)
}
