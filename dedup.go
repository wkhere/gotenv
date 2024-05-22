package gotenv

import (
	"cmp"
	"slices"
)

func dedup(a []EnvVar) []EnvVar {
	slices.SortStableFunc(a, func(e1, e2 EnvVar) int {
		return cmp.Compare(e1.Key, e2.Key)
	})

	for i := 0; i < len(a)-1; {
		if a[i].Key == a[i+1].Key {
			a = append(a[:i], a[i+1:]...)
		} else {
			i++
		}
	}
	return a
}
