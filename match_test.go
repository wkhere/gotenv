package gotenv

import (
	"reflect"
	"strings"
	"testing"
)

type tcase struct {
	input   string
	result  envvar
	wanterr string
}

func tvalid(input, name, val string) tcase {
	return tcase{input: input, result: envvar{name, val}}
}
func terror(input, errmsg string) tcase {
	return tcase{input: input, wanterr: errmsg}
}

var tcs = []tcase{
	tvalid("", "", ""),
	tvalid(" ", "", ""),
	tvalid("     ", "", ""),
	tvalid("\t", "", ""),
	tvalid("\t\t\t\t", "", ""),
	tvalid(" \t\t  \t ", "", ""),
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
	tvalid(`aa=x'x`, "aa", "x'x"),
	tvalid(`aa=x''x`, "aa", "x''x"),
	tvalid(`aa=x"x`, "aa", `x"x`),
	tvalid(`aa=x""x`, "aa", `x""x`),
	tvalid(`aa=x'"x`, "aa", `x'"x`),
	tvalid(`aa=x"'x`, "aa", `x"'x`),
	tvalid("_a=1", "_a", "1"),
	tvalid("a_=1", "a_", "1"),
	tvalid("a1=1", "a1", "1"),
	tvalid("_1=1", "_1", "1"),
	terror("1=1", "invalid state"),
	terror("1=", "invalid state"),
	tvalid("aa=12 ", "aa", "12"),
	tvalid(" aa=12 ", "aa", "12"),
	tvalid("aa =12", "aa", "12"),
	tvalid("aa= 12", "aa", "12"),
	tvalid("aa = 12", "aa", "12"),
	tvalid(" aa= 12 ", "aa", "12"),
	tvalid(" aa = 12 ", "aa", "12"),
	tvalid(" aa = '12' ", "aa", "12"),
	tvalid(` aa = "12" `, "aa", "12"),
	terror(`aa = 1 2`, "invalid state"),
	tvalid(`aa = 12#`, "aa", "12"),
	tvalid(`aa = 12# foo`, "aa", "12"),
	tvalid(`a=#12`, "a", ""),
	tvalid(`a=#`, "a", ""),
	tvalid(`a=1#2`, "a", "1"),
	tvalid(`a=12#`, "a", "12"),
	tvalid(`a=12 #`, "a", "12"),
	tvalid(`a=12# foo`, "a", "12"),
	tvalid(`a="12#"`, "a", "12#"),
	tvalid(`aa = "12"#`, "aa", "12"),
	tvalid(`aa = "12"# foo`, "aa", "12"),
	tvalid(`aa = 12 #`, "aa", "12"),
	tvalid(`aa = "12" #`, "aa", "12"),
	tvalid(`aa = "12" # foo`, "aa", "12"),
	tvalid(`# aa = "12"`, "", ""),
	tvalid(`#aa = "12"`, "", ""),
	tvalid(` #aa = "12"`, "", ""),
	tvalid(`    #aa = "12"`, "", ""),
	tvalid(" \t\t  #aa = '12'", "", ""),
	terror(`a#a="12"`, "invalid state"),
	terror(`a#= "12"`, "invalid state"),
	terror(`a# ="12"`, "invalid state"),
	tvalid(`a=#"12"`, "a", ""),
	tvalid(`a=# "12"`, "a", ""),
}

func TestMatch(t *testing.T) {
	for i, tc := range tcs {
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

func BenchmarkMatch(b *testing.B) {
	const cap = 40
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs[:cap] {
			match([]byte(tc.input))
		}
	}
}
