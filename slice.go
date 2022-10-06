package sliceutils

// Unique returns unique elements from slice.
func Unique[T comparable](s []T) []T {
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
func Filter[T any](s []T, f func(T) bool) []T {
	passed := make([]T, 0, len(s)/2)

	for i := range s {
		if f(s[i]) {
			passed = append(passed, s[i])
		}
	}

	return passed
}

// GroupBy return hashmap with values grouped by keys.
func GroupBy[K comparable, V any](s []V, f func(V) K) map[K][]V {
	groups := make(map[K][]V)

	for i := range s {
		group := f(s[i])
		groups[group] = append(groups[group], s[i])
	}

	return groups
}
