package sortarraybyparity

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"four elements",
			args{[]int{3, 1, 2, 4}},
			[]int{2, 4, 3, 1},
		},
		{
			"one element",
			args{[]int{0}},
			[]int{0},
		},
		{
			"empty",
			args{[]int{}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
