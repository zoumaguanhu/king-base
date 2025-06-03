package dict

import (
	"fmt"
	"king.com/king/base/common/strs"
	"testing"
)

func TestDict(t *testing.T) {
	c := &Currency{
		PrefixLab: "$",
		SuffixLab: "USD",
	}
	s := strs.ObjToStr(c)
	fmt.Printf("TestDict result:%v", s)
}
