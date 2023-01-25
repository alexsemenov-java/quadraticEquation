package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("Реши квадратное уравнение")

	fmt.Println("Введите a:")
	fmt.Scan(&a)

	fmt.Println("Введите b:")
	fmt.Scan(&b)

	fmt.Println("Введите c:")
	fmt.Scan(&c)

	D := (b * b) - 4*a*c
}
