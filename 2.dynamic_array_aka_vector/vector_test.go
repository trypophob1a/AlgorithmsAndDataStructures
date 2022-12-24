package vector

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size []int
	}
	type testCase[T any] struct {
		name    string
		args    args
		wantLen *Vector[T]
		wantErr bool
	}
	tests := []testCase[int]{
		{"too many arguments", args{[]int{1, 2, 3, 4, 5}}, &Vector[int]{}, true},
		{
			"has only capacity",
			args{[]int{0, 2}},
			&Vector[int]{0, 2, make([]int, 0, 2)}, false,
		},
		{
			"empty args",
			args{[]int{}},
			&Vector[int]{0, 10, make([]int, 0, 10)}, false,
		},
		{
			"has only len",
			args{[]int{3}},
			&Vector[int]{3, 3, make([]int, 3)}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New[int](tt.args.size...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantLen) {
				t.Errorf("New() got = %v, wantLen %v", got, tt.wantLen)
			}
		})
	}
}

func Test_vectorArgsHandler(t *testing.T) {
	type args struct {
		size []int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantCap int
		wantErr bool
	}{
		{"too many arguments", args{[]int{1, 2, 3, 4, 5}}, 0, 0, true},
		{"empty size", args{}, 0, 10, false},
		{"len is negative", args{[]int{-1}}, 0, 0, true},
		{"only one argument", args{[]int{3}}, 3, 3, false},
		{"capacity is negative", args{[]int{3, -1}}, 0, 0, true},
		{"is normal", args{[]int{1, 5}}, 1, 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := vectorArgsHandler(tt.args.size...)
			if (err != nil) != tt.wantErr {
				t.Errorf("vectorArgsHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantLen {
				t.Errorf("vectorArgsHandler() got = %v, wantLen %v", got, tt.wantLen)
			}
			if got1 != tt.wantCap {
				t.Errorf("vectorArgsHandler() got1 = %v, wantLen %v", got1, tt.wantCap)
			}
		})
	}
}

func TestVector_Len(t *testing.T) {
	type testCase[T any] struct {
		name string
		v    Vector[T]
		want int
	}
	tests := []testCase[int]{
		{"is empty", Vector[int]{len: 0, cap: 2}, 0},
		{"one element", Vector[int]{len: 1, cap: 2}, 1},
		{"five elements", Vector[int]{len: 5, cap: 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_grow(t *testing.T) {
	type args[T any] struct {
		item T
		rest []T
	}
	type testCase[T any] struct {
		name      string
		v         Vector[T]
		args      args[T]
		arrayWant []T
		lenWant   int
		capWant   int
	}

	tests := []testCase[int]{
		{
			"grow three elements",
			Vector[int]{2, 2, []int{1, 2}},
			args[int]{3, []int{4, 5}},
			[]int{1, 2, 3, 4, 5},
			5, 7,
		},
		{
			"grow one elements",
			Vector[int]{2, 2, []int{1, 2}},
			args[int]{3, []int{}},
			[]int{1, 2, 3},
			3, 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.grow(tt.args.item, tt.args.rest...)
			if got := len(tt.v.array); got != tt.lenWant {
				t.Errorf("Len = %v, lenWant %v", got, tt.lenWant)
			}

			if got := cap(tt.v.array); got != tt.capWant {
				t.Errorf("Cap = %v, capWant %v", got, tt.capWant)
			}

			if got := tt.v.array; !reflect.DeepEqual(got, tt.arrayWant) {
				t.Errorf("array = %v, arrayWant %v", got, tt.arrayWant)
			}
		})
	}
}

func TestVector_Append(t *testing.T) {
	type args[T any] struct {
		item T
		rest []T
	}
	type testCase[T any] struct {
		name      string
		v         *Vector[T]
		args      args[T]
		wantLen   int
		wantCap   int
		wantArray []T
	}
	tests := []testCase[int]{
		{
			"without grow", &Vector[int]{2, 5, []int{1, 2}},
			args[int]{3, []int{4}},
			4, 5,
			[]int{1, 2, 3, 4},
		},
		{
			"with grow capacity is greater than length", &Vector[int]{2, 5, []int{1, 2}},
			args[int]{3, []int{4, 5, 6}},
			6, 14,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"with grow capacity equal length", &Vector[int]{2, 2, []int{1, 2}},
			args[int]{3, []int{4, 5, 6}},
			6, 8,
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Append(tt.args.item, tt.args.rest...)
			if got := tt.v.len; got != tt.wantLen {
				t.Errorf("Len = %v, wantLen %v", got, tt.wantLen)
			}

			if got := tt.v.cap; got != tt.wantCap {
				t.Errorf("Cap = %v, wantCap %v", got, tt.wantCap)
			}

			if got := tt.v.array; !reflect.DeepEqual(got, tt.wantArray) {
				t.Errorf("array = %v, wantArray %v", got, tt.wantArray)
			}
		})
	}
}

func TestVector_String(t *testing.T) {
	t.Run("is empty", func(t *testing.T) {
		vec, _ := New[int]()
		if got := vec.String(); got != "[]" {
			t.Errorf("String() = %v, want %v", got, "[]")
		}
	})

	type args[T any] struct {
		item T
		rest []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{"one element", args[int]{item: 1}, "[1]"},
		{"three elements", args[int]{1, []int{2, 3}}, "[1 2 3]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec, _ := New[int]()
			vec.Append(tt.args.item, tt.args.rest...)

			if got := vec.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_removeByIndex(t *testing.T) {
	type args struct {
		index int
	}
	type testCase[T any] struct {
		name      string
		v         *Vector[T]
		args      args
		wantLen   int
		wantArray []T
		wantErr   bool
	}
	tests := []testCase[int]{
		{
			"is empty", &Vector[int]{0, 0, make([]int, 0)},
			args{0},
			0,
			[]int{},
			true,
		},
		{
			"one element", &Vector[int]{1, 2, []int{1}},
			args{0},
			0,
			[]int{},
			false,
		},
		{
			"two elements delete first", &Vector[int]{2, 2, []int{1, 2}},
			args{0},
			1,
			[]int{2},
			false,
		},
		{
			"two elements delete last", &Vector[int]{2, 2, []int{1, 2}},
			args{1},
			1,
			[]int{1},
			false,
		},
		{
			"three elements delete first", &Vector[int]{3, 10, []int{1, 2, 3}},
			args{0},
			2,
			[]int{2, 3},
			false,
		},
		{
			"three elements delete second", &Vector[int]{3, 10, []int{1, 2, 3}},
			args{1},
			2,
			[]int{1, 3},
			false,
		},
		{
			"three elements delete last", &Vector[int]{3, 10, []int{1, 2, 3}},
			args{2},
			2,
			[]int{1, 2},
			false,
		},
		{
			"three elements delete with error", &Vector[int]{3, 10, []int{1, 2, 3}},
			args{5},
			3,
			[]int{1, 2, 3},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.RemoveByIndex(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("removeByIndex() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.v.len != tt.wantLen {
				t.Errorf("removeByIndex() len = %v, wantLen %v", tt.v.len, tt.wantLen)
			}

			if got := tt.v.array; !reflect.DeepEqual(got, tt.wantArray) {
				t.Errorf("array = %v, wantArray %v", got, tt.wantArray)
			}
		})
	}
}

func TestVector_removeFirst(t *testing.T) {
	type testCase[T any] struct {
		name      string
		v         *Vector[T]
		wantLen   int
		wantArray []T
		wantErr   bool
	}
	tests := []testCase[int]{
		{
			"vector is empty", &Vector[int]{0, 0, make([]int, 0)},
			0,
			[]int{},
			true,
		},
		{
			"vector have one element", &Vector[int]{1, 1, []int{1}},
			0,
			[]int{},
			false,
		},
		{
			"vector have two element", &Vector[int]{2, 2, []int{1, 2}},
			1,
			[]int{2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.RemoveFirst(); (err != nil) != tt.wantErr {
				t.Errorf("removeFirst() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.v.len != tt.wantLen {
				t.Errorf("removeFirst() len = %v, wantLen %v", tt.v.len, tt.wantLen)
			}

			if got := tt.v.array; !reflect.DeepEqual(got, tt.wantArray) {
				t.Errorf("array = %v, wantArray %v", got, tt.wantArray)
			}
		})
	}
}

func TestVector_removeLast(t *testing.T) {
	type testCase[T any] struct {
		name      string
		v         *Vector[T]
		wantLen   int
		wantArray []T
		wantErr   bool
	}
	tests := []testCase[int]{
		{
			"vector is empty", &Vector[int]{0, 0, make([]int, 0)},
			0,
			[]int{},
			true,
		},
		{
			"vector have one element", &Vector[int]{1, 1, []int{1}},
			0,
			[]int{},
			false,
		},
		{
			"vector have two element", &Vector[int]{2, 2, []int{1, 2}},
			1,
			[]int{1},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.RemoveLast(); (err != nil) != tt.wantErr {
				t.Errorf("removeLast() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.v.len != tt.wantLen {
				t.Errorf("removeLast() len = %v, wantLen %v", tt.v.len, tt.wantLen)
			}

			if got := tt.v.array; !reflect.DeepEqual(got, tt.wantArray) {
				t.Errorf("array = %v, wantArray %v", got, tt.wantArray)
			}
		})
	}
}

func TestVector_InsertByIndex(t *testing.T) {
	type args[T any] struct {
		index int
		item  T
		rest  []T
	}
	type testCase[T any] struct {
		name      string
		v         *Vector[T]
		args      args[T]
		wantArray []T
		wantLen   int
		wantCap   int
	}
	tests := []testCase[int]{
		{
			"insert in empty vec with index greater than len without grow",
			&Vector[int]{0, 10, make([]int, 0, 10)},
			args[int]{1, 1, []int{}},
			[]int{1},
			1, 10,
		},
		{
			"insert in empty vec by index 0 without grow",
			&Vector[int]{0, 10, make([]int, 0, 10)},
			args[int]{0, 1, []int{}},
			[]int{1},
			1, 10,
		},
		{
			"insert 1 element by index 0 in not empty vec without grow",
			&Vector[int]{3, 10, []int{2, 3, 4}},
			args[int]{0, 1, []int{}},
			[]int{1, 2, 3, 4},
			4, 10,
		},
		{
			"insert 3 elements by index 1 in not empty vec without grow",
			&Vector[int]{3, 10, append(make([]int, 0, 10), 1, 5, 6)},
			args[int]{1, 2, []int{3, 4}},
			[]int{1, 2, 3, 4, 5, 6},
			6, 10,
		},
		{
			"insert 3 elements penultimate cell without grow",
			&Vector[int]{3, 10, append(make([]int, 0, 10), 1, 2, 6)},
			args[int]{2, 3, []int{4, 5}},
			[]int{1, 2, 3, 4, 5, 6},
			6, 10,
		},
		{
			"insert 3 elements first cell with grow",
			&Vector[int]{3, 3, append(make([]int, 0, 3), 4, 5, 6)},
			args[int]{0, 1, []int{2, 3}},
			[]int{1, 2, 3, 4, 5, 6},
			6, 9,
		},
		{
			"insert 3 elements by index 1 in not empty vec with grow",
			&Vector[int]{3, 3, append(make([]int, 0, 3), 1, 5, 6)},
			args[int]{1, 2, []int{3, 4}},
			[]int{1, 2, 3, 4, 5, 6},
			6, 9,
		},
		{
			"insert 3 elements penultimate cell with grow",
			&Vector[int]{3, 3, append(make([]int, 0, 3), 1, 2, 6)},
			args[int]{2, 3, []int{4, 5}},
			[]int{1, 2, 3, 4, 5, 6},
			6, 9,
		},
		{
			"insert 3 elements first cell with grow",
			&Vector[int]{3, 3, append(make([]int, 0, 3), 4, 5, 6)},
			args[int]{0, 1, []int{2, 3}},
			[]int{1, 2, 3, 4, 5, 6},
			6, 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.InsertByIndex(tt.args.index, tt.args.item, tt.args.rest...)
			if got := tt.v.array; !reflect.DeepEqual(got, tt.wantArray) {
				t.Errorf("array = %v, wantArray %v", got, tt.wantArray)

				if tt.v.len != tt.wantLen {
					t.Errorf("InsertByIndex() len = %v, wantLen %v", tt.v.len, tt.wantLen)
				}

				if tt.v.cap != tt.wantCap {
					t.Errorf("InsertByIndex() Cap = %v, wantCap %v", tt.v.cap, tt.wantCap)
				}
			}
		})
	}
}

func TestVector_Reverse(t *testing.T) {
	type testCase[T any] struct {
		name string
		v    Vector[T]
		want []T
	}
	tests := []testCase[int]{
		{
			"reverse five elements",
			Vector[int]{5, 5, append(make([]int, 0, 5), 1, 2, 3, 4, 5)},
			[]int{5, 4, 3, 2, 1},
		},
		{
			"reverse empty",
			Vector[int]{0, 0, []int{}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_GetVector(t *testing.T) {
	type testCase[T any] struct {
		name string
		v    Vector[T]
		want []T
	}
	tests := []testCase[int]{
		{
			"empty",
			Vector[int]{0, 0, []int{}},
			[]int{},
		},
		{
			"five elements",
			Vector[int]{5, 5, append(make([]int, 0, 5), 1, 2, 3, 4, 5)},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetVector(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Set(t *testing.T) {
	type args[T any] struct {
		index int
		item  T
	}
	type testCase[T any] struct {
		name    string
		v       Vector[T]
		args    args[T]
		wantArr []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			"vector is empty",
			Vector[int]{0, 0, make([]int, 0)},
			args[int]{0, 1},
			[]int{},
			true,
		},
		{
			"index out of range",
			Vector[int]{3, 3, append(make([]int, 0, 3), 1, 2, 3)},
			args[int]{3, 4},
			[]int{1, 2, 3},
			true,
		},
		{
			"good index",
			Vector[int]{3, 3, append(make([]int, 0, 3), 1, 2, 3)},
			args[int]{2, 4},
			[]int{1, 2, 4},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Set(tt.args.index, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := tt.v.GetVector(); !reflect.DeepEqual(got, tt.wantArr) {
				t.Errorf("array = %v, want %v", got, tt.wantArr)
			}
		})
	}
}

func TestVector_GetItem(t *testing.T) {
	type args struct {
		index int
	}
	type testCase[T any] struct {
		name    string
		v       Vector[T]
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			"vector is empty",
			Vector[int]{0, 0, make([]int, 0)},
			args{0},
			0,
			true,
		},
		{
			"index out of range",
			Vector[int]{3, 3, append(make([]int, 0, 3), 1, 2, 3)},
			args{3},
			0,
			true,
		},
		{
			"good index",
			Vector[int]{3, 3, append(make([]int, 0, 3), 1, 2, 3)},
			args{1},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.GetItem(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
