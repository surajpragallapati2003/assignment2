package main

import "fmt"

type bank struct {
	Name string
	acno uint
	ifsc string
	age  uint
}
type vehicle struct {
	model string
}

func (v *vehicle) check() {
	switch v.model {
	case "car":
		fmt.Println("it is a car")
	case "bike":
		fmt.Println("it is a bike")
	case "truck":
		fmt.Println("it is a truck")
	}
}

func (b *bank) String() string {
	return fmt.Sprintf("Name %v", b.Name)
	//fmt.Println(b.Name)
	//fmt.Println(b.acno)
	//fmt.Println(b.ifsc)
	//fmt.Println(b.age)
}

func greater(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return b
	}
	return c
}
func age(a uint) bool {
	if a >= 18 {
		return true
	} else {
		return false
	}
}
func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}
func calculater(a, b float32, c string) float32 {
	switch c {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func sum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s
}
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {

	if age(18) {
		fmt.Println("eligible for voting")
	} else {
		fmt.Println("not eligible for voting")
	}
	//fmt.Println(greater(3, 2, 1))
	//fmt.Println(isEven(6))
	//fmt.Println(isEven(7))
	//fmt.Println(calculater(3.1, 2.4, "*"))
	//fmt.Println(calculater(3.0, 2.4, "+"))
	//fmt.Println(calculater(3.0, 2.4, "*"))
	//fmt.Println(sum(10))
	//a := 10
	//b := 3
	//fmt.Println("a is %d \n b is %d", a, b)
	//swap(&a, &b)
	//fmt.Println("a is %d \n b is %d", a, b)
	//b := bank{"suraj", 1233446, "abcd1234", 21}
	//fmt.Println(b.Name)
	//fmt.Println(b.acno)
	//fmt.Println(b.ifsc)
	//fmt.Println(b.age)
	//b.print()
	vehicle := vehicle{"car"}
	vehicle.check()
	fmt.Println(&bank{Name: "Jay"})
}
