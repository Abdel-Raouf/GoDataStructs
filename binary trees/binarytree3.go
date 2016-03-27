package main

import (
  "fmt"
)



type node struct {
  left *node
  right *node
  isTree bool
  val int
}


func insertBST(head *node, newVal int) {
  newNode := &node{
    val: newVal,
  }
  if head.isTree == false{
    head.val = newVal
    head.isTree = true
    return
  }
  for{
    if newNode.val > head.val{
      if head.right == nil{
        head.right = newNode
        return
      }
      head = head.right
    }else if newNode.val <= head.val{
      if head.left == nil{
        head.left = newNode
        return
      }
      head = head.left
    }
  }
}

func printPath(head *node, val int, path *[]int) bool {
  if head.val == val{
    *path = append(*path, head.val)
    return true
  }
  if head.right != nil && printPath(head.right, val, path) == true{
      *path = append(*path, head.val)
      return true
  }
  if head.left != nil && printPath(head.left, val, path) == true{
      *path = append(*path, head.val)
      return true
  }
  return false
}

func main() {
  head := new(node)

  insertBST(head, 5)
  insertBST(head, 1)
  insertBST(head, 6)
  insertBST(head, 3)
  insertBST(head, 2)
  insertBST(head, 7)
  insertBST(head, 6)
  insertBST(head, 11)

  path := make([]int,0)

  printPath(head, 2, &path)

  for i := len(path)/2-1; i >= 0; i-- {
    opp := len(path)-1-i
    path[i], path[opp] = path[opp], path[i]
  }
  fmt.Println(path)

}
