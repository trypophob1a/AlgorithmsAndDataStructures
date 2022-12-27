package containsduplicateii

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"first and last",
			args{[]int{1, 2, 3, 1}, 3},
			true,
		},
		{
			"penultimate and last",
			args{[]int{1, 0, 1, 1}, 1},
			true,
		},
		{
			"dont have",
			args{[]int{1, 2, 3, 1, 2, 3}, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_naiveContainsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"first and last",
			args{[]int{1, 2, 3, 1}, 3},
			true,
		},
		{
			"penultimate and last",
			args{[]int{1, 0, 1, 1}, 1},
			true,
		},
		{
			"dont have",
			args{[]int{1, 2, 3, 1, 2, 3}, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := naiveContainsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("naiveContainsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
