package main

import (
  "fmt"
)

type node struct {
  left *node
  right *node
  val int
  isTree bool
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
    if newNode.val > head.val {
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

func getPath(myNode *node, myVal int, path *[]int) bool {
  if myNode.val == myVal {
    *path = append(*path, myNode.val)
    return true
  }
  if myNode.right != nil && getPath(myNode.right, myVal, path) == true{
    *path = append(*path, myNode.val)
    return true
  }
  if myNode.left != nil && getPath(myNode.left, myVal, path) == true{
    *path = append(*path, myNode.val)
    return true
  }
  return false
}


func main() {
  head := new(node)

  insertBST(head, 5)
  insertBST(head, 3)
  insertBST(head, 7)
  insertBST(head, 1)
  insertBST(head, 6)
  insertBST(head, 9)

  path := make([]int, 0)

  getPath(head, 1, &path)
  for i := len(path)/2-1; i >= 0; i-- {
    opp := len(path)-1-i
    path[i], path[opp] = path[opp], path[i]
  }
  fmt.Println(path)
}
