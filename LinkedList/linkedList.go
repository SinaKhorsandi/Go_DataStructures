package LinkedList

import "fmt"

type LinkedList struct {
	First *Node
	Last  *Node
}

type Node struct {
	Next *Node
	Data interface{}
}

func New() *LinkedList {
	emptyNode := &Node{
		Next: nil,
		Data: nil,
	}
	return &LinkedList{
		First: emptyNode,
		Last:  emptyNode,
	}
}

func (ll *LinkedList) Add(d interface{}) *LinkedList {
	NextNode:=&Node{
		Next: nil,
		Data: d,
	}
	if ll.First.Data == nil {
		ll.First = NextNode
	}else {
		ll.Last.Next=NextNode
	}
	ll.Last = NextNode
	return ll
}

func (ll *LinkedList) DeleteNode(d interface{}) *LinkedList {
	var node = ll.First
	if node.Data == d {
		node = node.Next
		return ll
	}
	for {
		if node.Next.Data == d {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
				break
			} else {
				node.Next = nil
				break
			}
		}
		node = node.Next
	}
	return ll
}


func (ll *LinkedList) IsEmpty() bool {
	if ll.First == nil {
		return true
	}
	return false
}

func (ll *LinkedList) PrintAll()  {
	var node = ll.First
	for  {
		fmt.Println(node.Data)
		if node.Next == nil {
			return
		}
		node = node.Next
	}
}
