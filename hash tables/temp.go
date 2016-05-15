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

func reSize(myTable *table) *table {
  newTable := createTable(myTable.length * 2)
  for i := range myTable.list{
    for tableColumn := myTable.list[i]; tableColumn != nil; tableColumn = tableColumn.next{
      insertTable(newTable, tableColumn.val)
    }
  }
  return newTable
}

func insertTable(myTable *table, myVal string) *table {
  if getLength(myTable.list) == (myTable.length - 1) {
    myTable = reSize(myTable)
  }else{
    location := findSpot(myTable.length, myVal)
    myTable.list[location] = insertNode(myTable.list[location], myVal)
  }
  return myTable
}

func main() {
  myTal := createTable(5)

  for _, word := range os.Args[1:] {
    myTal = insertTable(myTal, word)
  }
  
}
