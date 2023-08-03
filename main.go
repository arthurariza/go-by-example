package main

import (
	"fmt"
	"math"
)

func main() {
	// hello_world()

	// values()

	// variables()

	constant()
}

func hello_world() {
	fmt.Println("Hello World!")
}

func values()  {
	fmt.Println("Go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

const s string = "constant"

func constant() {
	fmt.Println(s)

	const n = 5000000
	const m = 500_000

	fmt.Println(n, m)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
