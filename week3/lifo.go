package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Stack struct {
 head *Node
}

func (st * Stack) insert(val int){
    temp := &Node{val : val, next: nil}
    if st.head == nil{
        st.head = temp
    } else{
        temp.next = st.head
        st.head = temp
    }
}
func (st *Stack) pop() *Node {
 temp := st.head
 if st.head != nil {
  if st.head.next == nil {
   st.head = nil
  } else {
   st.head = st.head.next
  }
 }
 return temp
}
func (st *Stack) peek() *Node {
 return st.head
}

func (st *Stack) clear() {
 st.head = nil
}

func (st *Stack) print() {
 cur := st.head
 for cur != nil {
  fmt.Println(cur.val)
  cur = cur.next
 }
}

func (st *Stack) increment(val int) {
 cur := st.head
 for cur != nil {
  //fmt.Println(cur.val)
  cur.val += val
  cur = cur.next
 }
}

func (st *Stack) printReverse() {
 ar := make([]int, 0)
 cur := st.head
 for cur != nil {
  fmt.Println(cur.val)
  ar = append(ar, cur.val)
  cur = cur.next
 }
 for i := 0; i < len(ar); i++ {
  fmt.Println(ar[i])
 }
}

func (st *Stack) contains(val int) bool {
 temp := st.head
 for temp != nil {
  if temp.val == val {
   return true
  }
 }
 return false
}

func main() {
 st := Stack{head: nil}
 st.insert(5)
 st.insert(4)
 st.pop()
 st.print()
}