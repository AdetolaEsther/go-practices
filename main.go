package main

import "fmt"

// packages and imports
// func main() {
// 	name := "Ada"
//     fmt.Println("Hello", name)
// }

// func main() {
//     fmt.Println("I am learning Go")
//     fmt.Println("I am testing imports")
// }

// unused import
// import (
//     "fmt"
//     "strings"
// )

// func main() {
//     fmt.Println("Hello, Go")
// }

// func main() {
//     name := "ada"
//     fmt.Println(strings.ToUpper(name))
// }

// Type Inference in Go
//   func main() {
// 	// var declares itemCount with the type int and the value 3.
// 	var itemCount int = 3
// 	// price := 2.5 declares a new variable. Go infers float64 from 2.5.
//     price := 29.5
// 	// name := "tea" declares a string. The := creates a new name here.
// 	name := "tea"
// 	// A plain = assigns to an existing variable. It does not declare a new one, so no := is needed.
// 	itemCount = itemCount + 6
// 	// Prints the three values in order: tea itemcount and price
// 	fmt.Println(name, itemCount, price)
// }

// Understanding a for loop
// func main() {
// 	// This happens once, before any check,The body runs first when the check is true, runs while i is 3 or less and adds 1 after each pass.
//     for i := 0; i < 3; i++ {
//         fmt.Println(i)
//     }
//     }

// Range Loops
// func main() {
//     for i := range 5 {
//         if i == 2 {
//             continue
//         }

//         fmt.Println(i)
//     }
// }

// Switch Statement

// Expression switch, evaluate the value, finds a match, execute the matching case and exit the switch
// func main() {
//     route := "/profgile"

//     switch route {
//     case "/":
//         fmt.Println("Home page")
//     case "/dashboard":
//         fmt.Println("Dashboard page")
//     case "/profile":
//         fmt.Println("Profile page")
//     default:
//         fmt.Println("Page not found")
//     }
// }
// Go also allows you to omit the value after switch. Each case can then be a Boolean condition
//  func main(){
// 	score := 75

// 	switch {
// 	case score >= 80:
// 		fmt.Println("Excellent")
// 	case score >= 60:
// 		fmt.Println("Passed")
// 	default:
// 		fmt.Println("Try again")
// 	}
//  }

//  Type switch
// func main() {
//     var value any = true

//     switch v := value.(type) {
//     case string:
//         fmt.Println("This is a string:", v)
//     case int:
//         fmt.Println("This is an integer:", v)
//     case bool:
//         fmt.Println("This is a Boolean:", v)
//     default:
//         fmt.Println("Unknown type")
//     }
// }

// Mixing incompatible types
//  func main(){
// 	age := 25

// switch age {
// case "25":
//     fmt.Println("Twenty-five")
// }
//  }

// Final practice exercise
// Create a small program that combines what you've learned.
// Your program should:
// 1. Declare a variable called status and assign it "pending".
// 2. Use an expression switch to print a message for "pending", "approved", or "rejected".
// 3. Print "Unknown status" for any other value.
// 4. Use a for loop with range 5 to print the numbers zero through four, but skip the number two using continue
// func main() {
//     status := "pending"

//     switch status {
//     case "pending":
//         fmt.Println("Pending approval")
//     case "approved":
//         fmt.Println("Approved!")
//     case "rejected":
//         fmt.Println("Rejected.")
//     default:
//         fmt.Println("Unknown status")
//     }
// 	 for i := range 5 {
//         if i == 2 {
//             continue
//         }

//         fmt.Println(i)
//     }
// }

//DATA STRUVT
//array & sliices

//  Declare an array without specifying initial values
// func main() {
//     // var numbers [3]int
// var names [3]string

//     // fmt.Println(numbers)
//     fmt.Println(names)
// }

// finding the length of an array with len
// func main() {
// 	// numbers := [4]int{10, 20, 30, 40}
// 	// length := len(numbers)
// 	// fmt.Println(length)
// 	// fmt.Println(numbers[3])
// 	numbers := [4]int{10, 20, 30, 40}

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}

// }
// assigning values to an array
// func main() {
// 	first := [3]int{10, 20, 30}
// 	second := first

// 	second[0] = 99

// 	fmt.Println(first)  // [10 20 30]
// 	fmt.Println(second) // [99 20 30]
// }

// SLICES
// func main() {
// 	numbers := [5]int{10, 20, 30, 40, 50}
// 	middle := numbers[0:4]
// 	fmt.Println(middle)
// 	fmt.Println(len(middle))
// }

// A. Declare a slice without initializing it
//
//	func main() {
//		var numbers []int
//		numbers = append(numbers, 10)
//		numbers = append(numbers, 20)
//		fmt.Println(numbers)
//		fmt.Println(len(numbers))
//		fmt.Println(numbers == nil)
//	}
//
// B. Initialize a slice with a slice literal
// func main() {
// 	numbers := []int{10, 20, 30}
// 	names := []string{"Ada", "Grace", "Linus"}
// 	active := []bool{true, false, true}
// 	prices := []float64{12.5, 18.75, 20.0}
// 	fmt.Println(numbers)
// 	fmt.Println(names)
// 	fmt.Println(active)
// 	fmt.Println(prices)
// }

// C. Create a slice using make
//
//	func main() {
//		numbers := make([]int, 3)
//		//You can then assign values by index:
//		numbers[0] = 10
//		numbers[1] = 20
//		numbers[2] = 30
//		fmt.Println(numbers)
//	}
//
// D. Create a slice with a specified capacity
// func main() {
// 	numbers := make([]int, 3, 5)
// 	numbers[2] = 20
// 	//To add a fourth element, use append:
// 	numbers = append(numbers, 40)
// 	fmt.Println(numbers)
// 	fmt.Println(cap(numbers))

// }
// Underlying arrays and shared data
// func main() {
// 	numbers := [5]int{10, 20, 30, 40, 50}

// 	middle := numbers[1:4]
// 	first := numbers[0:3]
// 	first[2] = 300
// 	second := numbers[2:5]
// 	middle[0] = 99
// 	fmt.Println(numbers)
// 	fmt.Println(first)
// 	fmt.Println(second)
// 	fmt.Println(middle)

// }

// length vs capacity
// func main() {
// 	numbers := make([]int, 2, 5)
// 	numbers[0] = 10
// 	numbers[1] = 20
// 	numbers = append(numbers, 30)
// 	numbers[2] = 30

// 	// numbers = append(numbers, 40)
// 	// numbers = append(numbers, 50)
// 	// numbers = append(numbers, 60)
// 	// numbers = append(numbers, 70)
// 	// numbers = append(numbers, 80)

// 	fmt.Println(numbers)
// 	fmt.Println(len(numbers))
// 	fmt.Println(cap(numbers))

// }

// Adding elements with append
// func main() {
// 	// numbers := []int{10, 20, 30}

// 	// // numbers = append(numbers, 40)
// 	// //You can append more than one element in a single operation:
// 	// numbers = append(numbers, 30, 40, 50)
// 	// first := []int{10, 20, 30}
// 	// second := []int{40, 50}
// 	// //You can append the elements of second to first using the ... operator:
// 	// first = append(first, second...)
// 	numbers := make([]int, 2, 5)
// 	numbers[0] = 10
// 	numbers[1] = 20

// 	numbers = append(numbers, 30)

// 	fmt.Println(numbers)
// 	// fmt.Println(first)
// 	// fmt.Println(second)
// }

// 2.7. Removing an element from a slice
// func main() {
// 	numbers := []int{10, 20, 30, 40, 50}

// 	i := 3

// 	numbers = append(numbers[:i], numbers[i+1:]...)

// 	fmt.Println(numbers)
// }

// // Reslicing
// func main() {
// 	numbers := []int{10, 20, 30, 40, 50}

// 	first := numbers[1:5]
// 	second := first[1:3]

// 	fmt.Println(first)

// 	fmt.Println(second)

// }

// 2.8. Copying slices
// func main() {
// 	first := []int{10, 20, 30}
// 	second := make([]int, len(first)) //creates a new slice with three elements. Then:

// 	copy(second, first) //copies the three values from first into second.

// 	second[0] = 99

// 	fmt.Println(first)
// 	fmt.Println(second)

// }

// 2.9. Nil slices versus empty slices
// func main() {
// 	var first []int
// 	second := []int{}
// 	fmt.Println(len(first))  // 0
// 	fmt.Println(len(second)) // 0

// 	fmt.Println(first == nil)  // true
// 	fmt.Println(second == nil) // false
// }

//MAPS

// Declaring and initializing maps
// func main() {
// 	userAges := map[string]int{
// 		"Adetola": 25,
// 		"Tobi":    30,
// 		"David":   28,
// 	}
// 	fmt.Println(userAges)

// }
// Using make()
// func main() {
// 	userAges := make(map[string]int)

// 	userAges["Adetola"] = 25
// 	userAges["Tobi"] = 30

// 	fmt.Println(userAges)
// }

// Declaring a nil map
// func main() {
// 	userAges := make(map[string]int)
// 	userAges["Adetola"] = 25 // panic

//		fmt.Println(userAges)
//	}
//
// Adding and updating values in a map
// func main() {
// 	userAges := make(map[string]int)

// 	userAges["Adetola"] = 25
// 	//updating
// 	userAges["Adetola"] = 26
// 	userAges["Tobi"] = 30

// 	fmt.Println(userAges)
// }

// 5. Accessing values in a map
// func main() {
// 	userAges := make(map[string]int)
// 	userAges["Adetola"] = 26

// 	age := userAges["Adetola"]

// 	fmt.Println(age) // 26

// }
// Checking whether a key exists
// func main() {
// 	userAges := make(map[string]int)
// 	userAges["Adetola"] = 26
// 	age, exists := userAges["Adetola"]

// 	if exists {
// 		fmt.Println("Age:", age)
// 	} else {
// 		fmt.Println("User not found")
// 	}

// }

//Structs: Creating struct variables (instances)
//Method 1: Declaring a struct variable using var

// type User struct {
//     Name  string
//     Age   int
//     Email string
// }

// func main() {
//     var user1 User

//	    fmt.Println(user1)
//	}
//
// zero values
// type User struct {
// 	Name     string
// 	Age      int
// 	Balance  float64
// 	IsActive bool
// }

// func main() {
// 	var user1 User

// 	fmt.Println(user1.Name)
// 	fmt.Println(user1.Age)
// 	fmt.Println(user1.Balance)
// 	fmt.Println(user1.IsActive)
// }

//Method 2: Using struct literals

// type User struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func main() {
// 	user1 := User{
// 		Name:  "Adetola",
// 		Age:   25,
// 		Email: "adetola@example.com",
// 	}

// 	fmt.Println(user1)
// }

// Method 3: Positional initialization
// type User struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func main() {
// 	user1 := User{"Adetola", 25, "adetola@example.com"}

// 	fmt.Println(user1)
// }

//Method 4: Declaring first and assigning values later

// type User struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func main() {
// 	var user1 User
// 	user1.Name = "Adetola"
// 	user1.Age = 25
// 	user1.Email = "adetola@example.com"

// 	fmt.Println(user1)
// }

//Modifying a struct field

// type User struct {
// 	Name    string
// 	Age     int
// 	Balance float64
// }

// func main() {
// 	user1 := User{
// 		Name:    "Adetola",
// 		Age:     25,
// 		Balance: 50000,
// 	}

// 	fmt.Println("Before update:", user1)

// 	user1.Balance = 75000

// 	fmt.Println("After update:", user1)
// }

//a slice of structs to store multiple users using name literal

//	type User struct {
//		Name string
//		Age  int
//	}
// type User struct {
// 	Name   string
// 	Scores []int
// }

// func main() {
// 	// users := []User{
// 	// 	{Name: "Adetola", Age: 25},
// 	// 	{Name: "Tobi", Age: 30},
// 	// 	{Name: "David", Age: 28},
// 	// }

// 	// fmt.Println(users)

// 	// //access individual fields using the dot operator
// 	// fmt.Println(users[0].Name)
// 	// fmt.Println(users[1].Age)
// 	user1 := User{
// 		Name:   "Adetola",
// 		Scores: []int{80, 90, 95},
// 	}

// 	user2 := user1

// 	user2.Scores[0] = 100

// 	fmt.Println(user1.Scores)
// 	fmt.Println(user2.Scores)
// }

//Exercise 1: Create a Student struct

// Create a struct called Student containing:
// - Name (string)
// - Age (int)
// - Grade (float64)
// - IsEnrolled (bool)
// Create two students and print their details.

// type Student struct {
// 	Name       string
// 	Age        int
// 	Grade      float64
// 	IsEnrolled bool
// }

// func main() {
// 	students := []Student{
// 		{Name: "ade", Age: 29, Grade: 82, IsEnrolled: true},
// 		{Name: "ade2", Age: 28, Grade: 89, IsEnrolled: true},
// 	}

// 	fmt.Println(students[0])
// 	fmt.Println(students[1])
// }

// Exercise 2: Modify struct fields

// Create a Product struct containing:
// - Name
// - Price
// - Quantity
// Create a product and update its price and quantity.

// Print the product before and after the update.

// type Product struct {
// 	Name     string
// 	Price    float64
// 	Quantity int
// }

// func main() {
// 	product1 := Product{
// 		Name:     "eba",
// 		Price:    5.00,
// 		Quantity: 6,
// 	}
// 	fmt.Println(product1.Price)
// 	product1.Price = 9.00
// 	fmt.Println(product1.Price)

// }
// Exercise 3: Slice of structs

// Create an Employee struct containing:
// - Name
// - Department
// - Salary
// Create a slice containing at least three employees.

// Use a for range loop to display each employee's information.

// type Employee struct {
// 	Name       string
// 	Department string
// 	Salary     int
// }

// func main() {
// 	Employee1 := Employee{
// 		Name:       "ayo",
// 		Department: "Tech",
// 		Salary:     4000000,
// 	}
// 	Employee2 := Employee{
// 		Name:       "uayo",
// 		Department: "ijhech",
// 		Salary:     40000,
// 	}
// 	Employee3 := Employee{
// 		Name:       "uayo",
// 		Department: "ijhech",
// 		Salary:     40000,
// 	}
// 	users := []Employee{Employee1, Employee2, Employee3}

// 	for index, user := range users {
// 		fmt.Println("User:", index+1)
// 		fmt.Println("Name:", user.Name)
// 		fmt.Println("Email:", user.Department)
// 		fmt.Println("Age:", user.Salary)
// 		fmt.Println("--------------------")
// 	}

// }

// func add(a, b int) int {
// 	result := a + b
// 	return result
// }
// func minus(a, b int) int {
// 	result := a - b
// 	return result
// }
// func multi(a, b int) int {
// 	result := a * b
// 	return result
// }

// func main() {
// 	total := add(10, 20)
// 	leftoverproduxt := minus(80, 40)
// 	toBuy := multi(2, 10)

// 	fmt.Println(total)
// 	fmt.Println("leftoverproduxt", leftoverproduxt)
// 	fmt.Println("leftoverproduxt", toBuy)
// }

// Returning Multiple Values

// func calculate(a, b int) (int, int) {
// 	sum := a + b
// 	difference := a - b

// 	return sum, difference
// }

// func main() {
// 	total, difference := calculate(20, 10)

// 	fmt.Println("Sum:", total)
// 	fmt.Println("Difference:", difference)
// }

// func add(numbers ...int) int {
// 	total := 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func main() {
// 	fmt.Println(add(10, 20))
// 	fmt.Println(add(10, 20, 30))
// 	fmt.Println(add(10, 20, 30, 40, 50))
// }

// func main() {
// 	age := 25

// 	agePointer := &age

// 	fmt.Println(agePointer)
// 	fmt.Println(*agePointer)
// }

// func main() {
// 	age := 25

// 	agePointer := &age

// 	*agePointer = 30

// 	fmt.Println(age)
// }

// nil pointers
// func main() {
// 	var agePointer *int

// 	fmt.Println(agePointer)
// }

// Pointers to structs
// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	user := User{
// 		Name: "Adetola",
// 		Age:  25,
// 	}

// 	userPointer := &user

// 	fmt.Println(userPointer)
// 	fmt.Println(*userPointer)
// }

//Exercise 1: Understanding addresses
// Create an integer variable called temperature with a value of 28.
// - Print its value.
// - Print its memory address.
// - Create a pointer to it.
// - Print the pointer's value.

// func main() {
// 	temperature := 28
// 	fmt.Println(temperature)
// 	fmt.Println(&temperature)
// 	temperaturePointer := temperature
// 	fmt.Println(temperaturePointer)
// }

// Exercise 2: Dereferencing
// Given:
// price := 5000
// pricePointer := &price

// Write code to:
// - Print the value through the pointer.
// - Change the value to 7500 using the pointer.
// - Print the original price variable.
// func main() {
// 	price := 5000
// 	pricePointer := &price
// 	*pricePointer = 7500
// 	fmt.Println(price)
// 	 fmt.Println(*pricePointer)

// }

// Exercise 3: Nil pointers
// Declare a pointer to a string.
// - Print its initial value.
// - Check whether it is nil.
// - Initialize it with the address of a string variable.
// - Print the value through the pointer.

// while creaying a nil pointer, short form declaration doesn’t work, you have to use var
// func main() {
// 	var isAwake *string
// 	fmt.Println(isAwake)
// 	if isAwake != nil {
// 		fmt.Println(*isAwake)
// 	} else {
// 		fmt.Println("Still sleeping")
// 	}
// 	status := "Awake"
// 	isAwake = &status
// 	fmt.Println(*isAwake)

// }

// Exercise 4: Pointers and structs
// Create a Product struct with:
// - Name (string)
// - Price (float64)
// - InStock (bool)
// Create a product, get its pointer, and use that pointer to update its price.
type Product struct {
	Name    string
	Price   float64
	Instock bool
}

func main() {
	product := Product{
		Name:    "chips",
		Price:   140.89,
		Instock: true,
	}
	productPointer := &product
	fmt.Println(productPointer)
	productPointer.Price = 250.9
	fmt.Println(product)
}
