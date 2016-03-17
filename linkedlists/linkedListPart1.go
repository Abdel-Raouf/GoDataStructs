package main

import(
	"fmt"
)

type Student struct{
	age int
	weight int
	name string
	Next * Student
}

type Teacher struct{
	age int
	weight int
	name string
	Next * Student
}


func main() {
	kyle := Student{7, 70, "Kyle", nil}
	john := Student{10, 100, "John", &kyle}

	debrah := Teacher{24, 185, "Debrah", &john}

	fmt.Println(debrah.name, john.name, kyle.name)

}
