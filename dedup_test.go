package gotenv

import (
	"reflect"
	"testing"
)

type ee = []EnvVar

var dedupTab = []struct {
	a, want []EnvVar
}{
	{a: nil, want: nil},
	{a: ee{}, want: ee{}},
	{a: ee{{Key: "foo", Val: "v1"}}, want: ee{{Key: "foo", Val: "v1"}}},
	{
		a: ee{
			{Key: "foo", Val: "v1"},
			{Key: "foo", Val: "v2"},
			{Key: "quux", Val: "v"},
		},
		want: ee{{Key: "foo", Val: "v2"}, {Key: "quux", Val: "v"}},
	},
	{
		a: ee{
			{Key: "quux", Val: "v"},
			{Key: "foo", Val: "v1"},
			{Key: "foo", Val: "v2"},
		},
		want: ee{{Key: "foo", Val: "v2"}, {Key: "quux", Val: "v"}},
	},
	{
		a: ee{
			{Key: "baz", Val: "b"},
			{Key: "foo", Val: "v1"},
			{Key: "foo", Val: "v2"},
		},
		want: ee{{Key: "baz", Val: "b"}, {Key: "foo", Val: "v2"}},
	},
	{
		a: ee{
			{Key: "baz", Val: "b1"},
			{Key: "baz", Val: "b2"},
			{Key: "baz", Val: "b3"},
			{Key: "foo", Val: "v1"},
			{Key: "foo", Val: "v2"},
		},
		want: ee{{Key: "baz", Val: "b3"}, {Key: "foo", Val: "v2"}},
	},
	{
		a: ee{
			{Key: "foo", Val: "v0"},
			{Key: "baz", Val: "b1"},
			{Key: "foo", Val: "v1"},
			{Key: "foo", Val: "v1.1"},
			{Key: "baz", Val: "b2"},
			{Key: "foo", Val: "v2"},
			{Key: "baz", Val: "b3"},
		},
		want: ee{{Key: "baz", Val: "b3"}, {Key: "foo", Val: "v2"}},
	},
}

func TestDedup(t *testing.T) {
	type ee = []EnvVar
	for i, tc := range dedupTab {
		a := dedup(tc.a)
		if !reflect.DeepEqual(a, tc.want) {
			t.Errorf("tc#%d mismatch\nhave: %v\nwant: %v", i, a, tc.want)
		}
	}
}

func BenchmarkDedup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range dedupTab[:8] {
			_ = dedup(tc.a)
		}
	}
}
