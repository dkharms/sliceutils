package sliceutils

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type testCase[T comparable] struct {
		input    []T
		expected []T
	}

	cases := []testCase[int]{
		{
			input:    nil,
			expected: []int{},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 1, 2, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, testCase := range cases {
		if actual := Unique(testCase.input); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}

func TestFilter(t *testing.T) {
	type testCase[T comparable] struct {
		input    []T
		f        func(T) bool
		expected []T
	}

	cases := []testCase[int]{
		{
			input: []int{1, 2, 3, 4},
			f: func(x int) bool {
				return x%2 != 0
			},
			expected: []int{1, 3},
		},
		{
			input: []int{1, 2, 3, 4},
			f: func(x int) bool {
				return x%2 == 0
			},
			expected: []int{2, 4},
		},
	}

	for _, testCase := range cases {
		if actual := Filter(testCase.input, testCase.f); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}

func TestGroupBy(t *testing.T) {
	type testCase[K comparable, V any] struct {
		input    []V
		f        func(V) K
		expected map[K][]V
	}

	cases := []testCase[string, int]{
		{
			input: []int{1, 2, 3, 4},
			f: func(x int) string {
				if x%2 != 0 {
					return "odd"
				}
				return "even"
			},
			expected: map[string][]int{
				"odd":  {1, 3},
				"even": {2, 4},
			},
		},
	}

	for _, testCase := range cases {
		if actual := GroupBy(testCase.input, testCase.f); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}

func TestIntersection(t *testing.T) {
	type testCase[T comparable] struct {
		x, y     []T
		expected []T
	}

	cases := []testCase[int]{
		{
			x:        []int{4, 5, 6},
			y:        []int{1, 1, 2, 3, 3},
			expected: []int{},
		},
		{
			x:        []int{1, 2, 3},
			y:        []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, testCase := range cases {
		if actual := Intersect(testCase.x, testCase.y); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}
