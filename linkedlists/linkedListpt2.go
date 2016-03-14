package main

import "fmt"


type Student struct{
	id int
	Next * Student
}

func reverse(curr * Student) * Student {
	if curr.Next == nil{
		return curr
	}else{
		temp := reverse(curr.Next)
		curr.Next.Next = curr
		curr.Next = nil
		return temp
	}
}

func main() {

	count := 5

	head := new(Student)
	head.Next, head.id = nil, count

	for index := 1; index < count; index++{
		temp := new(Student)

		temp.id = index

		temp.Next = head.Next
		head.Next = temp
	}

	fmt.Println("Decreasing Order: ")
	for temp := head; temp != nil; temp = temp.Next{
			fmt.Println(temp.id)
	}


	head = reverse(head)

	fmt.Println("Increasing Order: ")
	for ; head != nil; head = head.Next{
			fmt.Println(head.id)
	}

}
