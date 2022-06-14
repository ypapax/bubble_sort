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
			act := Sort(c.inp)
			if !as.Equal(c.exp, act) {
				return
			}
		})

	}

}
