package LinkedList

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
	if ll.First == nil {
		ll.First = NextNode
		return ll
	}else {
		ll.Last.Next=NextNode
		return ll
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
		if node.Next == d {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
				break
			} else {
				node.Next = nil
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
