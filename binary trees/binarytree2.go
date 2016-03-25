package main

import (
	"fmt"
	"math/rand"
)

const numbOfvals = 1000000
const medianOfvals = numbOfvals / 2

type bnode struct {
	left  *bnode
	right *bnode
	val   int
}

type llnode struct {
	next *llnode
	val  int
}

func binsert(head *bnode, val int) {
	temp := head
	newNode := &bnode{
		val: val,
	}
	for {
		if newNode.val < temp.val {
			if temp.left == nil {
				temp.left = newNode
				return
			}
			temp = temp.left
		}
		if newNode.val >= temp.val {
			if temp.right == nil {
				temp.right = newNode
				return
			}
			temp = temp.right
		}
	}
}


func main() {

	llhead := new(llnode)
	bhead := new(bnode)
	llhead.val = medianOfvals
  llhead.next = nil
	bhead.val = medianOfvals

	for i := 0; i < numbOfvals; i++ {
		tempRand := rand.Int()
		binsert(bhead, tempRand)

    temp := new(llnode)
    temp.val = tempRand
    temp.next = llhead
    llhead = temp

	}
  fmt.Println(llhead.next.val, bhead.right.left.val)
}
