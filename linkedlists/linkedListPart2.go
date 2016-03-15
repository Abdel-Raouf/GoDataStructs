package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

type student struct {
	name string
	age int
	ssn string
	next * student
}

func reverse(curr * student) * student {
	if curr.next == nil{
		return curr
	}else{
		last := reverse(curr.next)
		curr.next.next = curr
		curr.next = nil
		return last
	}
}


func main() {
	students := new(student)
	students.next = nil

  if len(os.Args) < 2{
    fmt.Println("There is no file")
    return
  }

	filename := os.Args[1]


	data, err := ioutil.ReadFile(filename)
	if err != nil{
		return
	}
	for _, line := range strings.Split(string(data), "\n") {

		if line == "" {
			continue
		}else{
			td := strings.Split(string(line), " ")

			ts := new(student)
			ts.name = td[0]
			ts.age, err = strconv.Atoi(td[1])
			ts.ssn = td[2]

			ts.next = students.next
			students.next = ts
		}

	}
	fmt.Println("Original----")
	for s := students.next; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}

	students = reverse(students)
	fmt.Println(" ")
	fmt.Println("Reversed----")
	for s := students; s.next != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}

}
