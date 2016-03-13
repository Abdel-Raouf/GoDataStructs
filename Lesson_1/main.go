package main

import "fmt"

type Student struct {
	Age int
	Weight int
	Name string
	Next * Vertex
}

func main() {
	count := 100


	tempStudent := Student{20, 160, "Kyle", nil}
	head := &tempStudent

	for NodeCount := 0; NodeCount < count; NodeCount++ {


		tempStudent = Student{NodeCount, NodeCount + 100, string(NodeCount) }


	}





}
