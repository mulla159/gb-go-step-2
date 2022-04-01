package sort

import (
	"reflect"
	"testing"
)

var globalResult = []int{}

func TestSortByInserts(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 6, 8, 7, 7}

	expected := []int{2, 4, 5, 6, 7, 7, 8, 9}

	SortByInserts(arr)

	if result := reflect.DeepEqual(arr, expected); result == false {
		t.Errorf("expected: %v\ngot %v", expected, arr)
	}
}

func TestBubleSort(t *testing.T) {
	var arr = []int{2, 4, 5, 9, 6, 8, 7, 7}

	expected := []int{2, 4, 5, 6, 7, 7, 8, 9}

	BubbleSort(arr)

	if result := reflect.DeepEqual(arr, expected); result == false {
		t.Errorf("expected: %v\ngot %v", expected, arr)
	}
}

func BenchBubbleSort(b *testing.B) {
	var arr = []int{2, 4, 5, 9, 6, 8, 7, 7}

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}

	globalResult = arr
}

func BenchSortByInserts(b *testing.B) {
	var arr = []int{2, 4, 5, 9, 6, 8, 7, 7}

	for i := 0; i < b.N; i++ {
		SortByInserts(arr)
	}

	globalResult = arr
}
