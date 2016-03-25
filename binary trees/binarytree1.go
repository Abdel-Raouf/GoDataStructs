package main

import (
  "fmt"
)

type node struct {
  left *node
  right *node
  val int
}

func insert(head *node, val int)  {
  newNode := new(node)
  newNode.val = val
  temp := head

  for {

    if val < temp.val {
      if temp.left == nil{
        temp.left = newNode
        return
      }
      temp = temp.left
    }else if val >= temp.val{
      if temp.right == nil{
        temp.right = newNode
        return
      }
      temp = temp.right
    }
  }
}


func main() {

  var head *node
  head = new(node)
  head.val = 1
  insert(head, 2)
  insert(head, 2)



  fmt.Println(head.val)
  fmt.Println(head.right.val)
  fmt.Println(head.right.right.val)
  for i := 0; i  < 100; i++ {
    insert(head, i)
  }

  fmt.Println(head.right.right.right.right.right.val)


}
