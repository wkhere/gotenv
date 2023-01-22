package gotenv

import (
	"reflect"
	"strings"
	"testing"
)

func TestMatch(t *testing.T) {
	type tcase struct {
		input   string
		result  envvar
		wanterr string
	}
	tvalid := func(input, name, val string) tcase {
		return tcase{input: input, result: envvar{name, val}}
	}
	terror := func(input, errmsg string) tcase {
		return tcase{input: input, wanterr: errmsg}
	}
	var tab = []tcase{
		terror("a", "invalid state"),
		terror("aaa", "invalid state"),
		terror("=", "invalid state"),
		terror("=aa", "invalid state"),
		tvalid("a=", "a", ""),
		tvalid("aa=1", "aa", "1"),
		tvalid("a=11", "a", "11"),
		tvalid("aa=1", "aa", "1"),
		tvalid("aaa=11", "aaa", "11"),
		tvalid(`aaa='11'`, "aaa", "11"),
		tvalid(`aaa="11"`, "aaa", "11"),
		terror(`aa='22`, "invalid state"),
		terror(`aa=22'`, "invalid state"),
		terror(`aa="22`, "invalid state"),
		terror(`aa=22"`, "invalid state"),
		terror(`aa='22"`, "invalid state"),
		terror(`aa="22'`, "invalid state"),
		tvalid("_a=1", "_a", "1"),
		tvalid("a_=1", "a_", "1"),
		tvalid("a1=1", "a1", "1"),
		tvalid("_1=1", "_1", "1"),
		terror("1=1", "invalid state"),
		terror("1=", "invalid state"),
	}

	for i, tc := range tab {
		e, err := match([]byte(tc.input))
		switch {
		case err != nil && tc.wanterr == "":
			t.Errorf("tc#%d: unexpected error: %v", i, err)
		case err == nil && tc.wanterr != "":
			t.Errorf("tc#%d: have no error while wanted one with msg: %s",
				i, tc.wanterr)
		case err != nil && tc.wanterr != "":
			if !strings.Contains(err.Error(), tc.wanterr) {
				t.Errorf(
					"tc#%d: have error `%v` while wanted one with submsg: %s",
					i, err, tc.wanterr,
				)
			}
		default:
			if !reflect.DeepEqual(e, tc.result) {
				t.Errorf("tc#%d mismatch:\nhave %v\nwant %v", i, e, tc.result)
			}
		}
	}
}
