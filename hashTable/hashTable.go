package hashTable

import (
	"github.com/dorin131/go-data-structures/linkedlist"
)

const arrayLength = 10

type HashTable struct {
	Data [arrayLength]*linkedlist.LinkedList
}

type DataList struct {
	Key   string
	value interface{}
}

func NewHash() *HashTable {
	return &HashTable{
		Data: [arrayLength]*linkedlist.LinkedList{},
	}
}

func hashFunc(k string) int {
	return len(k) % 10
}

func (h *HashTable) Add(k string, v interface{}) *HashTable {
	index := hashFunc(k)

	if h.Data[index] == nil {
		h.Data[index] = linkedlist.New()
		h.Data[index].Append(DataList{Key: k, value: v})
	} else {
		node := h.Data[index].Head
		for {
			if node != nil {
				d := node.Data.(DataList)
				if d.Key == k {
					d.value = v
					break
				}
			} else {
				h.Data[index].Append(DataList{
					k, v,
				})
				break
			}
			node = node.Next
		}
	}
	return h
}

func (h *HashTable) Get(k string) (v interface{}) {
	index := hashFunc(k)

	if h.Data[index] == nil {
		return ""
	}
		node:=h.Data[index].Head
		for  {
			d:=node.Data.(DataList)
			if d.Key == k {
				return v
			}else {
				return ""
			}
			node=node.Next
		}
}
