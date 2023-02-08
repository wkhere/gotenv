package gotenv

import (
	"reflect"
	"testing"
)

func TestDedup(t *testing.T) {
	type ee = []envvar
	var tab = []struct {
		a, want []envvar
	}{
		{a: nil, want: nil},
		{a: ee{}, want: ee{}},
		{a: ee{{key: "foo", val: "v1"}}, want: ee{{key: "foo", val: "v1"}}},
		{
			a: ee{
				{key: "foo", val: "v1"},
				{key: "foo", val: "v2"},
				{key: "quux", val: "v"},
			},
			want: ee{{key: "foo", val: "v2"}, {key: "quux", val: "v"}},
		},
		{
			a: ee{
				{key: "quux", val: "v"},
				{key: "foo", val: "v1"},
				{key: "foo", val: "v2"},
			},
			want: ee{{key: "foo", val: "v2"}, {key: "quux", val: "v"}},
		},
		{
			a: ee{
				{key: "baz", val: "b"},
				{key: "foo", val: "v1"},
				{key: "foo", val: "v2"},
			},
			want: ee{{key: "baz", val: "b"}, {key: "foo", val: "v2"}},
		},
		{
			a: ee{
				{key: "baz", val: "b1"},
				{key: "baz", val: "b2"},
				{key: "baz", val: "b3"},
				{key: "foo", val: "v1"},
				{key: "foo", val: "v2"},
			},
			want: ee{{key: "baz", val: "b3"}, {key: "foo", val: "v2"}},
		},
		{
			a: ee{
				{key: "foo", val: "v0"},
				{key: "baz", val: "b1"},
				{key: "foo", val: "v1"},
				{key: "foo", val: "v1.1"},
				{key: "baz", val: "b2"},
				{key: "foo", val: "v2"},
				{key: "baz", val: "b3"},
			},
			want: ee{{key: "baz", val: "b3"}, {key: "foo", val: "v2"}},
		},
	}

	for i, tc := range tab {
		a := dedup(tc.a)
		if !reflect.DeepEqual(a, tc.want) {
			t.Errorf("tc#%d mismatch\nhave: %v\nwant: %v", i, a, tc.want)
		}
	}
}
