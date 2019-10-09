package sat

import (
	"fmt"
	"testing"
)

func TestSat2(t *testing.T) {
	c := []clause{
		{6, 6},
		{-1, 2},
		{-1, -4},
		{1, 4},
		{-2, -3},
		{-3, -5},
		{-6, -5},
	}
	r, err := Sat2(6, c)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
