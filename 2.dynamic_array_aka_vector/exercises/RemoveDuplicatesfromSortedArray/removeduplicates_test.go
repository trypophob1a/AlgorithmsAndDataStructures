package removeduplicatesfromsortedarray

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantArr []int
	}{
		{"empty", args{[]int{}}, 0, []int{}},
		{"one element", args{[]int{1}}, 1, []int{1}},
		{
			"two uniq elements and arr len = 2",
			args{[]int{1, 2}},
			2,
			[]int{1, 2},
		},
		{
			"two uniq elements and arr len = 4",
			args{[]int{1, 2, 2, 3}},
			3,
			[]int{1, 2, 3},
		},
		{
			"one uniq element and arr len = 3",
			args{[]int{1, 1, 2}},
			2,
			[]int{1, 2},
		},
		{
			"five uniq element and arr len = 10",

			args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			5,
			[]int{0, 1, 2, 3, 4},
		},
		{
			"two uniq element and arr len = 5",

			args{[]int{0, 0, 0, 1, 1, 2, 2, 3, 3, 4}},
			5,
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicatesNaive() = %v, want %v", got, tt.want)
			}
			for i := 0; i < got; i++ {
				if tt.args.nums[i] != tt.wantArr[i] {
					t.Errorf("array element not equal  got %v, want %v", tt.args.nums[i], tt.wantArr[i])
				}
			}
		})
	}
}
