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
		if actual := Unique(testCase.input...); !reflect.DeepEqual(actual, testCase.expected) {
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
		if actual := Filter(testCase.f, testCase.input...); !reflect.DeepEqual(actual, testCase.expected) {
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
		if actual := GroupBy(testCase.f, testCase.input...); !reflect.DeepEqual(actual, testCase.expected) {
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

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		x        []T
		y        T
		expected bool
	}

	cases := []testCase[int]{
		{
			x:        []int{4, 5, 6},
			y:        7,
			expected: false,
		},
		{
			x:        []int{1, 2, 3},
			y:        3,
			expected: true,
		},
	}

	for _, testCase := range cases {
		if actual := Contains(testCase.y, testCase.x...); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}

func TestFindIndex(t *testing.T) {
	type testCase[T comparable] struct {
		x        []T
		y        T
		expected int
	}

	cases := []testCase[int]{
		{
			x:        []int{4, 5, 6},
			y:        7,
			expected: -1,
		},
		{
			x:        []int{1, 2, 3},
			y:        3,
			expected: 2,
		},
	}

	for _, testCase := range cases {
		if actual := FindIndex(testCase.y, testCase.x...); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}

func TestReverse(t *testing.T) {
	type testCase[T comparable] struct {
		x        []T
		expected []T
	}

	cases := []testCase[int]{
		{
			x:        []int{4, 5, 6},
			expected: []int{6, 5, 4},
		},
	}

	for _, testCase := range cases {
		if actual := Reverse(testCase.x...); !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, actual)
		}
	}
}

func TestMap(t *testing.T) {
	type testCase[T comparable] struct {
		x        []T
		f        func(T) T
		expected []T
	}

	cases := []testCase[int]{
		{
			x: []int{4, 5, 6},
			f: func(x int) int {
				return x
			},
			expected: []int{4, 5, 6},
		},
	}

	for _, testCase := range cases {
		if Map(testCase.f, testCase.x...); !reflect.DeepEqual(testCase.x, testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.x)
		}
	}
}
