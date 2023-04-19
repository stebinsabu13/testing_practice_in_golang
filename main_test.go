package check

import (
	"fmt"
	"testing"
)

type condition struct {
	arg1, arg2, want int
}

func TestMultiply(t *testing.T) {
	c := []condition{
		{4, 6, 24},
		{1, 2, 2},
		{5, 4, 20},
		{10, 11, 110},
	}
	for _, v := range c {
		t.Run(fmt.Sprintf("%v*%v", v.arg1, v.arg2), func(t *testing.T) {
			got := multiply(v.arg1, v.arg2)
			if got != v.want {
				t.Errorf("Expected %v but got %v", v.want, got)
			}
		})
	}
}
