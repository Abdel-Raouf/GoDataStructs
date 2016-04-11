package main

import (
  "fmt"
  "os"
)

type node struct {
  next *node
  val string
}

type table struct {
  list []*node
  length int
}

func createTable(newLen int) *table {
  return &table{
    list: make([]*node, newLen),
    length : newLen,
  }
}

func findSpot(length int, myVal string) int {
    var insertLoc int
    for _, i := range myVal {
      insertLoc += int(i)
    }
    return (insertLoc % length)
}

func insertNode(myNode *node, myVal string) *node {
  newNode := &node{
    next : nil,
    val : myVal,
  }
  if myNode != nil{
    newNode.next = myNode
  }
  return newNode
}

func getLength(vals []*node) int {
  var count int
  for _, pointer := range vals {
    if pointer != nil{
      count++
    }
  }
  return count
}

func insertTable(myTable *table, myVal string)  {
  if getLength(myTable.list) == (myTable.length - 1) {
    fmt.Println("we've filled the list") // TODO: resize table here
  }else{
    location := findSpot(myTable.length, myVal)
    myTable.list[location] = insertNode(myTable.list[location], myVal)
  }
  return
}

func main() {
  myTal := createTable(5)

  for _, word := range os.Args[1:] {
    insertTable(myTal, word)
  }
  fmt.Println(myTal)

}
