package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {

	var name string = "Eliodd"
	var age int = 20
	var isCool bool = true
	var thirdname string
	cake := "chocolate"
	sisterage := 23
	fmt.Println(name, age, isCool, thirdname, cake, sisterage)

	//memory address
	var children int16 = 12
	var amount float64 = 12.34
	expenses := 1.6
	fmt.Println(children, amount, expenses)
	fmt.Println("....")
	//printing and formatting strings
	fmt.Print("hello,kampala ")
	fmt.Print("hello,kenya")
	fmt.Println("....")
	emyaka := 6
	ezina := "Eliod"

	fmt.Println("my name is ", ezina, "and  i am ", emyaka, "years old")
	//printf formatted strings
	fmt.Printf("my name is %v and my age is %v \n", ezina, emyaka)
	fmt.Printf("my name is of type %T", ezina)

	var str = fmt.Sprintf("my name is %v and my age is %v \n", ezina, emyaka)
	fmt.Println("the value of str is ", str)
	//arrays and slices
	var fruits [2]string = [2]string{"mango", "orange"}
	fmt.Println(fruits)
	fruits[1] = "apples"
	fmt.Println(fruits, len(fruits))

	fmt.Println("....")

	var numbers = []int{1, 2, 6, 9, 6}
	fmt.Println(numbers, len(numbers))
	var numbers2 = append(numbers, 7)
	fmt.Println(numbers2, len(numbers2))

	//range
	rangeOne := numbers[1:]
	fmt.Println(rangeOne)
	//usingstrings
	greeting := "hello world"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(".......")
	ages := []int{12, 34, 45, 67, 34, 45}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 34)
	fmt.Println(index)

	names := []string{"Eliod", "muhangi", "tash", "Dyson", "andrew", "rickie"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println("...........")
	//forloop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//whileloop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	//range
	for index, value := range names {
		fmt.Printf("the index is %v and the value is %v \n", index, value)
	}
	//boolean and conditionals
	ags := 50
	fmt.Println(ags >= 500)
	fmt.Println(ags >= 34)
	fmt.Println("......................................................................................................")
	students := []string{"Eliod", "muhangi", "tash", "Dyson", "andrew", "rickie"}
	for index, value := range students {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Println("the value at pos", index, "is", value)
	}

	fmt.Println("......................................................................................................")
	fmt.Println(add(10, 24))
	fmt.Println("......................................................................................................")
	sayGoodBye("Eliod")
	fmt.Println("......................................................................................................")
	sayhello("Eliod")
	fmt.Println("......................................................................................................")
	cycleNames([]string{"joan", "jane", "jolly", "joseph", "josephat", "josephine"}, sayhello)
	fmt.Println("......................................................................................................")
	cycleNames([]string{"joan", "jane", "jolly", "joseph", "josephat", "josephine"}, sayGoodBye)
	fmt.Println("......................................................................................................")
	fmt.Println(areaOfCircle(12.5))
	fmt.Println(areaOfCircle(160.5))
	fmt.Println("......................................................................................................")
	fn, sn := getInitials("Eliod Muhangi")
	fmt.Println(fn, sn)
	fmt.Println("......................................................................................................")
	fmt.Println(addfournumbers(34, 45, 56, 67, 78))
	fmt.Println("......................................................................................................")
	fmt.Println(sayHi("Eliod"))
	fmt.Println("......................................................................................................")
	showname("Eliod")
	fmt.Println("......................................................................................................")

	///Golang maps
	menu := map[string]float64{
		"soup":    4.99,
		"pie":     6.78,
		"pizza":   56.0,
		"panakes": 67.0,
	}

	fmt.Println(menu)
	fmt.Println(menu["pizza"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//maps in int f
	phonenumbers := map[int]string{
		1234567890: "Eliod",
		1234567891: "Muhangi",
		2345678901: "Tash",
		3456789012: "Dyson",
	}
	phonenumbers[1234567890] = "mutiba"
	fmt.Println(phonenumbers)

	for k, v := range phonenumbers {
		fmt.Println(k, "-", v)
	}
	//passbyvalue

	name1 := "Elioda"
	name = updateName(name1)
	fmt.Println(name)

	//maps second type of vlaues
	phonenumbers1 := map[int]string{
		1234567890: "Eliod",
		1234567891: "Muhangi",
		2345678901: "Tash",
	}
	updatephonenumbers(phonenumbers1)
	fmt.Println(phonenumbers1)

	laptop := "macbook"
	fmt.Println("mememory addresss", &laptop)
	fmt.Println("......................................................................................................")
	m := &laptop
	fmt.Println("mememory addresss", m)
	fmt.Println("value at memory address", *m)
	fmt.Println(laptop)

	//pointers
	food := "rice"
	fmt.Println("memory address::", &food)
	foodie := &food
	fmt.Println("real word before change::", food)
	fmt.Println("pointer befor::", foodie)
	updateFood(foodie)
	fmt.Println("pointer after::", foodie)
	fmt.Println("real word after change::", food)
}

// passin pointers in as parameters
func updateFood(food *string) {
	*food = "pizza"
}

//pass by reference ..!

func updatephonenumbers(x map[int]string) {
	x[1234567890] = "mutiba"
	x[43537836489] = "joseph"
}

// passin by value
func updateName(x string) string {
	x = "wedge"
	return x
}

//pacakage scope

// SELF MADE FUNCTIONS
// functions that addss
func add(num1 int, num2 int) int {
	return num1 + num2
}
func sayhello(name string) {
	fmt.Println("hello", name)
}

// function that is used to great
func sayGoodBye(name string) {
	fmt.Println("GoodBye", name)
}

// function that allows function to be passed
func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

// area of the circle
func areaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

// a function that has many return types
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
