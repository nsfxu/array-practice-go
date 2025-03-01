package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	var hobbies [3]string = [3]string{"Learning", "Video games", "Programming"}

	// Task 1
	fmt.Println(hobbies)

	// Task 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// Task 3
	hobbiesTaskThree := hobbies[:2]
	// hobbiesTaskThree := hobbies[0:2]

	fmt.Println(hobbiesTaskThree)

	// Task 4
	hobbiesTaskFour := hobbiesTaskThree[1:]
	hobbiesTaskFour = append(hobbiesTaskFour, hobbies[2])

	fmt.Println(hobbiesTaskFour)

	// Task 5
	courseGoals := []string{"Learn GO", "Create an API with GO"}
	fmt.Println(courseGoals)

	// Task 6
	courseGoals[1] = "Get better at arrays"
	courseGoals = append(courseGoals, "Do more exercises with GO")

	fmt.Println(courseGoals)

	// Task 7

	productList := []Product{{id: "1", title: "Cellphone", price: 99.99}, {id: "2", title: "Remote Control", price: 25.32}}
	productList = append(productList, Product{id: "3", title: "Battery", price: 5})

	fmt.Println(productList)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
