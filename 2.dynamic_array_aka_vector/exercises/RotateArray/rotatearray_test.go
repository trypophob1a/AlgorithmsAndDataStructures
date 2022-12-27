package rotatearray

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name      string
		args      args
		wantArray []int
	}{
		{
			"1 rotate",
			args{[]int{-1, -100, 3, 99}, 1},
			[]int{99, -1, -100, 3},
		},
		{
			"2 rotate",
			args{[]int{-1, -100, 3, 99}, 2},
			[]int{3, 99, -1, -100},
		},
		{
			"3 rotate",
			args{[]int{-1, -100, 3, 99}, 3},
			[]int{-100, 3, 99, -1},
		},
		{
			"3 rotate if len = 3",
			args{[]int{1, 2, 3}, 3},
			[]int{1, 2, 3},
		},
		{
			"4 rotate if len = 3",
			args{[]int{1, 2, 3}, 4},
			[]int{3, 1, 2},
		},
		{
			"12 rotate if len = 3",
			args{[]int{1, 2, 3}, 3},
			[]int{1, 2, 3},
		},
		{
			"25 rotate if len = 3",
			args{[]int{1, 2, 3}, 3},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.wantArray) {
				t.Fatalf("got %v want %v", tt.args.nums, tt.wantArray)
			}
		})
	}
}
