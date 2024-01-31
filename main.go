package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// mybill := newBill("Muhangi")
	// mybill.updateTip(400)
	// fmt.Println(".............................")
	// mybill.addItem("mango", 4000)
	// mybill.addItem("rolex", 9000)
	// mybill.addItem("chapati", 5000)
	// mybill.addItem("soda", 2000)
	// mybill.addItem("water", 1000)

	// fmt.Println("............................")
	// fmt.Println(mybill.format())

	//input
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
	///interfaces
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill::")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill::", reader)
	b := newBill(name)
	fmt.Println("Created the bill::", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip):", reader)
	switch opt {
	case "a":
		name, _ := getInput("Enter item name:", reader)
		price, _ := getInput("Enter item price:", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added", name, price)
		promptOptions(b)
	case "s":
		tip, _ := getInput("Enter the tip", reader)
		p, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.addItem(tip, p)
		fmt.Println("item added", tip)
		promptOptions(b)
	case "t":
		b.save()
		fmt.Println("You have saved the file", b.name)
	default:
		fmt.Println("You choosen a wrong letter")
		promptOptions(b)
	}
}

// //interfacess
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}
