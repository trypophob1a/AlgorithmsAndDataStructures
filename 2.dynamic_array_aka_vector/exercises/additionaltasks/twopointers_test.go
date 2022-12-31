package additionaltasks

import (
	"reflect"
	"testing"
)

func Test_findPair(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"is empty", args{[]int{}, 5}, []int{}},
		{"len < 2", args{[]int{1}, 2}, []int{}},
		{"has two elements", args{[]int{1, 2, 3, 4, 5}, 7}, []int{2, 5}},
		{"the target is too large", args{[]int{1, 2, 3, 4, 5}, 10}, []int{}},
		{
			"first & mid index == sum",
			args{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, -5},
			[]int{-5, 0},
		},
		{
			"first & last index == sum",
			args{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, 0},
			[]int{-5, 5},
		},
		{
			"the target is too small",
			args{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, -10},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPair(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"result in right", args{[]int{1, 2, 3, 4, 5, 6, 7}, 5}, 4},
		{"result in mid", args{[]int{1, 2, 3, 4, 5, 6, 7}, 4}, 3},
		{"result in left", args{[]int{1, 2, 3, 4, 5, 6, 7}, 1}, 0},
		{"not found", args{[]int{1, 2, 3, 4, 5, 6, 7}, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.args.nums, tt.args.target)
			if got != tt.want {
				t.Errorf("binarySearch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPairWithBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"is empty", args{[]int{}, 5}, []int{}},
		{"len < 2", args{[]int{1}, 2}, []int{}},
		{"has two elements", args{[]int{1, 2, 3, 4, 5}, 7}, []int{2, 5}},
		{"the target is too large", args{[]int{1, 2, 3, 4, 5}, 10}, []int{}},
		{
			"first & mid index == sum",
			args{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, -5},
			[]int{-5, 0},
		},
		{
			"first & last index == sum",
			args{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, 0},
			[]int{-5, 5},
		},
		{
			"the target is too small",
			args{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, -10},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairWithBinarySearch(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPairWithBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindPair(b *testing.B) {
	b.Run("two pointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			findPair([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, -5)
			findPair([]int{1, 2, 3, 4, 5}, 10)
			findPair([]int{}, 5)
			findPair([]int{1, 2, 3, 4, 5}, 7)
		}
	})

	b.Run("binary search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			findPairWithBinarySearch([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, -5)
			findPairWithBinarySearch([]int{1, 2, 3, 4, 5}, 10)
			findPairWithBinarySearch([]int{}, 5)
			findPairWithBinarySearch([]int{1, 2, 3, 4, 5}, 7)
		}
	})
}
