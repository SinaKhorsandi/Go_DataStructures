package hashTable

import (
	"github.com/sinakhorsandi/testify/assert"
	"testing"
)

func TestHashTAble(t *testing.T) {
	a:=assert.New(t)
	d:=DataList{
		Key:"sina",value:"khorsandi" }

	h:=NewHash()
	nh:=h.Set(d.Key,d.value)
	v:=nh.Get(d.Key)
	a.Equal("khorsandi",v)

}