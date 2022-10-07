package sliceutils

// Unique returns unique elements from slice.
func Unique[T comparable](s ...T) []T {
	var (
		unique = make([]T, 0, len(s)/2)
		seen   = make(map[T]struct{})
	)

	for i := range s {
		if _, ok := seen[s[i]]; !ok {
			unique = append(unique, s[i])
		}
		seen[s[i]] = struct{}{}
	}

	return unique
}

// Filter returns values from slice which satisfies given predicate.
func Filter[T any](f func(T) bool, s ...T) []T {
	passed := make([]T, 0, len(s)/2)

	for i := range s {
		if f(s[i]) {
			passed = append(passed, s[i])
		}
	}

	return passed
}

// GroupBy return hashmap with values grouped by keys.
func GroupBy[K comparable, V any](f func(V) K, s ...V) map[K][]V {
	groups := make(map[K][]V)

	for i := range s {
		group := f(s[i])
		groups[group] = append(groups[group], s[i])
	}

	return groups
}

// Intersect returns intersection of two slices.
//
// If element appears several times in both slices then
// this element will be in resulted slice as many times as presented
// in slice with minimal element appearance.
func Intersect[T comparable](x, y []T) []T {
	var (
		seen         = make(map[T]int)
		intersection = make([]T, 0)
	)

	for i := range x {
		seen[x[i]]++
	}

	for i := range y {
		if counter, ok := seen[y[i]]; ok && counter > 0 {
			intersection = append(intersection, y[i])
			seen[y[i]]--
		}
	}

	return intersection
}

// Contains returns true if slice contains element.
func Contains[T comparable](v T, s ...T) bool {
	for i := range s {
		if v == s[i] {
			return true
		}
	}

	return false
}

// FindIndex return index of element in slice.
// If not found returns -1.
func FindIndex[T comparable](v T, s ...T) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

// Reverse reverses slice in-place.
func Reverse[T any](s ...T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// Map modifies each element of slice in-place.
func Map[T any](f func(T) T, s ...T) {
	for i := range s {
		s[i] = f(s[i])
	}
}
