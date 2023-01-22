
//line match.rl:1
package gotenv

import (
	"fmt"
)

func match(data []byte) (e envvar, err error) {
	
//line match.rl:9
	
//line match.go:14
var _envvar_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 3, 4, 4, 4, 4, 
}

const envvar_start int = 1
const envvar_first_final int = 8
const envvar_error int = 0

const envvar_en_main int = 1


//line match.rl:10

	cs, p, pe := 0, 0, len(data)
	eof := pe
	pb := 0

	text := func() string { return string(data[pb:p]) }
	textQuoted := func() string { return string(data[pb+1:p-1]) }

	
//line match.go:37
	{
	cs = envvar_start
	}

//line match.go:42
	{
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 1:
		if data[p] == 95 {
			goto tr0;
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0;
			}
		case data[p] >= 65:
			goto tr0;
		}
		goto tr1;
	case 0:
		goto _out
	case 2:
		switch data[p] {
		case 61:
			goto tr3;
		case 95:
			goto tr2;
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2;
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr2;
			}
		default:
			goto tr2;
		}
		goto tr1;
	case 8:
		switch data[p] {
		case 34:
			goto tr13;
		case 39:
			goto tr14;
		}
		goto tr12;
	case 9:
		switch data[p] {
		case 34:
			goto tr5;
		case 39:
			goto tr5;
		}
		goto tr4;
	case 3:
		switch data[p] {
		case 34:
			goto tr5;
		case 39:
			goto tr5;
		}
		goto tr4;
	case 4:
		if data[p] == 34 {
			goto tr7;
		}
		goto tr6;
	case 5:
		if data[p] == 34 {
			goto tr8;
		}
		goto tr6;
	case 10:
		if data[p] == 34 {
			goto tr8;
		}
		goto tr6;
	case 11:
		goto tr6;
	case 6:
		if data[p] == 39 {
			goto tr10;
		}
		goto tr9;
	case 7:
		if data[p] == 39 {
			goto tr11;
		}
		goto tr9;
	case 12:
		if data[p] == 39 {
			goto tr11;
		}
		goto tr9;
	case 13:
		goto tr9;
	}

	tr1: cs = 0; goto _again
	tr2: cs = 2; goto _again
	tr0: cs = 2; goto f0
	tr5: cs = 3; goto _again
	tr13: cs = 4; goto f0
	tr6: cs = 5; goto _again
	tr14: cs = 6; goto f0
	tr9: cs = 7; goto _again
	tr3: cs = 8; goto f1
	tr4: cs = 9; goto _again
	tr12: cs = 9; goto f0
	tr8: cs = 10; goto _again
	tr7: cs = 11; goto _again
	tr11: cs = 12; goto _again
	tr10: cs = 13; goto _again

f0:
//line match.rl:19
 pb = p 
	goto _again
f1:
//line match.rl:21
 e.key = text() 
	goto _again

_again:
	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		switch _envvar_eof_actions[cs] {
		case 3:
//line match.rl:22
 e.val = text() 
		case 4:
//line match.rl:23
 e.val = textQuoted() 
//line match.go:188
		}
	}

	_out: {}
	}

//line match.rl:35


	if cs < envvar_first_final {
		return e, fmt.Errorf("invalid state: %d", cs)
	}
	return e, nil
}
