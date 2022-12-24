package vector

import (
	"fmt"
)

type Vector[T any] struct {
	len, cap int
	array    []T
}

type IndexOutOfRange struct {
	error
}

func indexOutOfRange(index int) IndexOutOfRange {
	return IndexOutOfRange{fmt.Errorf("index[%d] out of range", index)}
}

func New[T any](size ...int) (*Vector[T], error) {
	length, capacity, err := vectorArgsHandler(size...)
	if err != nil {
		return &Vector[T]{}, err
	}

	return &Vector[T]{length, capacity, make([]T, length, capacity)}, nil
}

func vectorArgsHandler(size ...int) (int, int, error) {
	if len(size) > 2 {
		return 0, 0, fmt.Errorf("too many arguments %d. Should not exceed two", len(size))
	}

	if len(size) == 0 {
		return 0, 10, nil
	}

	if len(size) < 2 {
		if size[0] < 0 {
			return 0, 0, fmt.Errorf("len: %d can't be negative", size[0])
		}
		return size[0], size[0], nil
	}

	if len(size) == 2 && size[1] < 0 {
		return 0, 0, fmt.Errorf("capacity: %d can't be negative", size[1])
	}

	return size[0], size[1], nil
}

func (v *Vector[T]) Append(item T, rest ...T) {
	if v.len+len(rest)+1 > v.cap {
		v.grow(item, rest...)
	} else {
		v.array = append(v.array, item)
		v.array = append(v.array, rest...)
	}

	v.len += len(rest) + 1
}

func (v *Vector[T]) InsertByIndex(index int, item T, rest ...T) {
	if index > v.len-1 {
		v.Append(item, rest...)
		return
	}

	if v.len+len(rest)+1 > v.cap {
		v.cap = v.cap*2 + len(rest) + 1
		newArray := make([]T, 0, v.cap)
		newArray = append(newArray, v.array[:index]...)
		newArray = append(newArray, item)
		newArray = append(newArray, rest...)
		newArray = append(newArray, v.array[index:]...)
		v.array = newArray
	} else {
		v.array = append(v.array[:index+1+len(rest)], v.array[index:]...)
		v.array[index] = item
		for i := 0; i < len(rest); i++ {
			v.array[index+i+1] = rest[i]
		}
	}

	v.len += len(rest) + 1
}

func (v *Vector[T]) grow(item T, rest ...T) {
	v.cap = v.cap*2 + len(rest) + 1
	newArray := make([]T, len(v.array), v.cap)
	copy(newArray, v.array)
	newArray = append(newArray, item)
	newArray = append(newArray, rest...)
	v.array = newArray
}

func (v *Vector[T]) RemoveByIndex(index int) error {
	if v.len-1 < 0 || index > v.len-1 {
		return indexOutOfRange(index)
	}

	v.array = append(v.array[:index], v.array[index+1:]...)
	v.len--
	return nil
}

func (v *Vector[T]) RemoveFirst() error {
	if err := v.RemoveByIndex(0); err != nil {
		return err
	}
	return nil
}

func (v *Vector[T]) RemoveLast() error {
	if err := v.RemoveByIndex(v.len - 1); err != nil {
		return err
	}
	return nil
}

func (v *Vector[T]) Reverse() []T {
	newVec := append(make([]T, 0, v.cap), v.array...)
	last := len(newVec)
	for i := 0; i < last/2; i++ {
		newVec[last-1-i], newVec[i] = newVec[i], newVec[last-1-i]
	}
	return newVec
}

func (v *Vector[T]) Set(index int, item T) error {
	if index > v.len-1 {
		return indexOutOfRange(index)
	}

	v.array[index] = item
	return nil
}

func (v *Vector[T]) GetItem(index int) (T, error) {
	if index > v.len-1 {
		var item T
		return item, indexOutOfRange(index)
	}

	return v.array[index], nil
}

func (v *Vector[T]) GetVector() []T {
	return v.array
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", v.array)
}

func (v *Vector[T]) Len() int {
	return v.len
}
