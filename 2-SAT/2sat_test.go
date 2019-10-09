package sat

import (
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
	for _, i := range c {
		b1, b2 := false, false
		if i.x < 0 {
			b1 = !r[-i.x]
		} else {
			b1 = r[i.x]
		}

		if i.y < 0 {
			b2 = !r[-i.y]
		} else {
			b2 = r[i.y]
		}
		if !(b1 || b2) {
			t.Fatal("all clauses should be true")
		}
	}
}
