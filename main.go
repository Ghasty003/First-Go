package main

import "fmt"

type Person struct {
	name string
	age  int
}

func checkAge(age int) (result bool) {
	if age >= 18 {
		result = true
		return
	} else {
		result = false
		return
	}
}

func iterateArr() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func iterateSlice() {
	slice := []string{"John", "Doe", "Maz", "Ilya"}

	for _, i := range slice {
		fmt.Println(i)
	}
}

func sayWelcome() {
	fmt.Println("Welcome!")
}

func main() {

	var p Person

	p.age = 20
	p.name = "Ghasty"

	fmt.Println(p.age, p.name)

	iterateSlice()
	iterateArr()

	age := 2

	if checkAge(age) {
		sayWelcome()
		return
	}

	fmt.Println("You need to be 18 or higher")
}
