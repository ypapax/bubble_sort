package bubble_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	type testCase struct {
		inp, exp []int
	}
	cases := []testCase{
		{inp: []int{4, 2, 1}, exp: []int{1, 2, 4}},
		{inp: []int{1, 2, 4}, exp: []int{1, 2, 4}},
		{inp: []int{1, 88, 2, 99, 4}, exp: []int{1, 2, 4, 88, 99}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			as := assert.New(t)

			act := BubbleSort(c.inp)
			t.Logf("actual: %+v", act)
			if !as.Equal(c.exp, act) {
				return
			}

			act2 := SelectionSort(c.inp)
			t.Logf("actual2: %+v", act2)
			if !as.Equal(c.exp, act2) {
				return
			}
		})
	}
}
