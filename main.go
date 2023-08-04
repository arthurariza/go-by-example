package main

import (
	"fmt"
	"math"
)

func main() {
	// hello_world()

	// values()

	// variables()

	// constant()

	// forr()

	IfElse()
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

func forr() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func IfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
			fmt.Println(num, "has 1 digit")
	} else {
			fmt.Println(num, "has multiple digits")
	}
 
	// No Ternary Operator Go 
	// 7%2 == 0 ? fmt.Println("7 is even") : fmt.Println("7 is odd")
}
