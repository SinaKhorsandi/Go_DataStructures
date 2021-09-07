package LinkedList

import (
	"github.com/sinakhorsandi/testify/assert"
	"testing"
)


func ExampleLinkedList_Add() {
	ll := New()
	ll.Add("first data").Add("second data").PrintAll()
	// Output:
	// first data
	// second data
}
func ExampleLinkedList_DeleteNode() {
	ll := New()
	ll.Add("first data").Add("second data").DeleteNode("second data").PrintAll()
	// Output:
	// first data
}

func TestLinkedList_IsEmpty(t *testing.T) {
	a:=assert.New(t)
	ll:= New()
	b:=ll.Add("second data").Add("third data").DeleteNode("second data").IsEmpty()
	a.Equal(false,b)
}
