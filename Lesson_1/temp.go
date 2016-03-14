package main

import(
	"fmt"
)

type Student struct{
	Age int
	Weight int
	Name string
	Next * Student
}

type Teacher struct{
	Age int
	Weight int
	Name string
	Subject string
	Next * Student
}



func main() {
	kyle := Student{20, 165, "Kyle", nil}
	john := Student{21, 200, "John", &kyle}

	debrah := Teacher{35, 155, "Debrah", "Computer Science", &john}

	fmt.Println("The teacher is:", debrah.Name)
	fmt.Println("Who is standing right behind her?")
	fmt.Println(debrah.Next.Name, "is standing behind her")




}
