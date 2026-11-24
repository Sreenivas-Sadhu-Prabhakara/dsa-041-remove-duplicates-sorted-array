package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{7, 7, 7}, []int{7}},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
	}
	for _, c := range cases {
		in := append([]int{}, c.in...)
		n := RemoveDuplicates(in)
		if !reflect.DeepEqual(in[:n], c.want) {
			t.Errorf("RemoveDuplicates(%v) kept %v, want %v", c.in, in[:n], c.want)
		}
	}
}
