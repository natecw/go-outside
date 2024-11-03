package main

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := New(func(i1, i2 int) bool { return i1 < i2 })
	ints := []int{2, 11, 8, 4, 9, 6}
	want := []int{2, 4, 6, 8, 9, 11}

	for _, v := range ints {
		heap.Push(v)
	}

	for i := range ints {
		t.Run(fmt.Sprintf("pop(%d)", i), func(t *testing.T) {
			got := heap.Pop()
			if want[i] != got {
				t.Errorf("got %d, want %d\n", got, want[i])
			}
		})
	}
}

func TestMaxHeap(t *testing.T) {
	heap := New(func(i1, i2 int) bool { return i1 > i2 })
	ints := []int{2, 11, 8, 4, 9, 6}
	want := []int{11, 9, 8, 6, 4, 2}

	for _, v := range ints {
		heap.Push(v)
	}

	for i := range ints {
		t.Run(fmt.Sprintf("pop(%d)", i), func(t *testing.T) {
			got := heap.Pop()
			if want[i] != got {
				t.Errorf("got %d, want %d\n", got, want[i])
			}
		})
	}
}
