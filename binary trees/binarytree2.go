package main

import (
	"fmt"
	"math/rand"
	"time"
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


func findBST(head *bnode, val int) (time.Duration, string, int){
	start := time.Now()
	comparisons := 0
	for {
		if val < head.val {
			comparisons++
			if head.left == nil {
				return time.Since(start), "not found", comparisons
			}
			head = head.left
		}else if val > head.val {
			comparisons++
			if head.right == nil {
				return time.Since(start), "not found", comparisons
			}
			head = head.right
		}else{
			break
		}
	}
	return time.Since(start), "found", comparisons
}

func findLinkedList(head *llnode, val int) (time.Duration, string, int) {
	start := time.Now()
	comparisons := 0
	for s := head; s != nil; s = s.next{
		comparisons++
		if s.val == val{
			return time.Since(start), "found", comparisons
		}
	}
	return time.Since(start), "not found", comparisons
}


func main() {

	llhead := new(llnode)
	bhead := new(bnode)
	llhead.val = medianOfvals
  llhead.next = nil
	bhead.val = medianOfvals

	for i := 0; i < numbOfvals; i++ {
		tempRand := rand.Int() % 100000
		binsert(bhead, tempRand)

		temp := &llnode{
			val: tempRand,
			next:llhead,
		}
		llhead = temp
	}

	BSTtime, found, comparisons := findBST(bhead, 6969)
	fmt.Println("Time to find on BST:", BSTtime, "and it was", found, comparisons, "comparisons")
	LinkedListTime, found, comparisons  := findLinkedList(llhead, 6969)
	fmt.Println("Time to find on Linked List:", LinkedListTime, "and it was", found,"at", comparisons, "comparisons")
}
