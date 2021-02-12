package algorithm

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	cases := []struct {
		i, j     int
		in, want []int
	}{
		{1, 3, []int{1, 2, 2, 1}, []int{1, 1, 2, 2}},
		{0, 1, []int{1, 2, 2, 1}, []int{2, 1, 2, 1}},
	}

	for _, c := range cases {
		pre := make([]int, len(c.in))
		copy(pre, c.in)
		Swap(c.in, c.i, c.j)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Swap in: %v == %v, expected: %v", pre, c.in, c.want)
		} else {
			t.Logf("Swap in: %v == %v, expected: %v", pre, c.in, c.want)
		}
	}
}

func TestQSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{
			[]int{2, 3, 4, 1, 2},
			[]int{1, 2, 2, 3, 4},
		}, {
			[]int{},
			[]int{},
		}, {
			[]int{1},
			[]int{1},
		}, {
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, c := range cases {
		pre := make([]int, len(c.in))
		copy(pre, c.in)
		got := QSort(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("QSort in: %v == %v, expected: %v", pre, got, c.want)
		} else {
			t.Logf("QSort in: %v == %v, expected: %v", pre, got, c.want)
		}
	}

}
