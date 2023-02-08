package gotenv

import "golang.org/x/exp/slices"

func dedup(a []envvar) []envvar {
	slices.SortStableFunc(a, func(e1, e2 envvar) bool {
		return e1.key < e2.key
	})

	for i := 0; i < len(a)-1; {
		if a[i].key == a[i+1].key {
			a = append(a[:i], a[i+1:]...)
		} else {
			i++
		}
	}
	return a
}
