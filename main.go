package main

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
import (
	"fmt"
)

const Pi = 3.14

func double(x int) int {
	return x * 2
}
func greet(name string) string {
	return "Hello " + name
}
func perimeter(rad float64) float64 {
	return 2 * Pi * rad
}
func area(rad float64) float64 {
	return rad * rad * Pi
}
func main() {
	fmt.Println(double(2))
	fmt.Println(greet("Suraj"))
	fmt.Println(perimeter(2.5))
	fmt.Println(area(2.5))

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
