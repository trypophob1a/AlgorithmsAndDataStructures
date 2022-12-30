package threesum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"is empty", args{[]int{1, -1}}, [][]int{}},
		{"two elements in slice", args{[]int{1, -1}}, [][]int{}},
		{"has one triplet & len = 3", args{[]int{-1, 0, 1}}, [][]int{{-1, 0, 1}}},
		{
			"has two triplets & len = 6",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"has three triplets & len = 9",
			args{[]int{-1, 0, 1, 2, -1, -4, 4, 0, 15}},
			[][]int{{-4, 0, 4}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"has one triplet all value equal zero",
			args{[]int{0, 0, 0}},
			[][]int{{0, 0, 0}},
		},
		{
			"has one triplet four zero",
			args{[]int{0, 0, 0, 0}},
			[][]int{{0, 0, 0}},
		},
		{
			"has one triplet  & len = 4",
			args{[]int{1, -1, -1, 0}},
			[][]int{{-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
