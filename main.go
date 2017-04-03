package main

import (
	"./binary"
	"./dijsktra"
	"./fibonacci"
	"./palindrome"
	"./palindrome2"
	"./print_star"
	"./reverse"
	"./up"
	"fmt"
)

func main() {
	fmt.Println("test package binary")
	binary.Generate()
	fmt.Println("test package dijsktra")
	dijsktra.Generate()
	fmt.Println("test package fibonacci")
	fibonacci.Generate()
	fmt.Println("test package palindrome")
	palindrome.Generate()
	fmt.Println("test package palindrome2")
	palindrome2.Generate()
	fmt.Println("test package print_star")
	print_star.Generate()
	fmt.Println("test package reverse")
	reverse.Generate()
	fmt.Println("test package up")
	up.Generate()
}
