package hashTable

import (
	"github.com/dorin131/go-data-structures/linkedlist"
	"math"
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
		[arrayLength]*linkedlist.LinkedList{},
	}
}

func hashFunc(s string) int {
	h := 0
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func index(hash int) int {
	return hash % arrayLength
}
func (h *HashTable) Set(k string, v interface{}) *HashTable {
	index := index(hashFunc(k))

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
	index := index(hashFunc(k))
	if h.Data[index] == nil {
		return ""
	}
	node := h.Data[index].Head
	for {
		if node != nil {
			d := node.Data.(DataList)
			if d.Key == k {
				return d.value
			} else {
				return ""
			}
			node = node.Next
		}
	}
}
