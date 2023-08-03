package main

import "fmt"

func main() {
	hello_world()

	values()
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
