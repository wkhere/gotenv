
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
	0, 0, 3, 4, 4, 4, 4, 
}

const envvar_start int = 7
const envvar_first_final int = 7
const envvar_error int = 0

const envvar_en_main int = 7


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
	case 7:
		switch data[p] {
		case 32:
			goto tr11;
		case 95:
			goto tr12;
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr11;
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr12;
			}
		default:
			goto tr12;
		}
		goto tr1;
	case 0:
		goto _out
	case 8:
		if data[p] == 32 {
			goto tr11;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr11;
		}
		goto tr1;
	case 1:
		switch data[p] {
		case 61:
			goto tr2;
		case 95:
			goto tr0;
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr0;
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr0;
			}
		default:
			goto tr0;
		}
		goto tr1;
	case 9:
		switch data[p] {
		case 34:
			goto tr14;
		case 39:
			goto tr15;
		}
		goto tr13;
	case 10:
		switch data[p] {
		case 34:
			goto tr4;
		case 39:
			goto tr4;
		}
		goto tr3;
	case 2:
		switch data[p] {
		case 34:
			goto tr4;
		case 39:
			goto tr4;
		}
		goto tr3;
	case 3:
		if data[p] == 34 {
			goto tr6;
		}
		goto tr5;
	case 4:
		if data[p] == 34 {
			goto tr7;
		}
		goto tr5;
	case 11:
		if data[p] == 34 {
			goto tr7;
		}
		goto tr5;
	case 12:
		goto tr5;
	case 5:
		if data[p] == 39 {
			goto tr9;
		}
		goto tr8;
	case 6:
		if data[p] == 39 {
			goto tr10;
		}
		goto tr8;
	case 13:
		if data[p] == 39 {
			goto tr10;
		}
		goto tr8;
	case 14:
		goto tr8;
	}

	tr1: cs = 0; goto _again
	tr0: cs = 1; goto _again
	tr12: cs = 1; goto f1
	tr4: cs = 2; goto _again
	tr14: cs = 3; goto f1
	tr5: cs = 4; goto _again
	tr15: cs = 5; goto f1
	tr8: cs = 6; goto _again
	tr11: cs = 8; goto _again
	tr2: cs = 9; goto f0
	tr3: cs = 10; goto _again
	tr13: cs = 10; goto f1
	tr7: cs = 11; goto _again
	tr6: cs = 12; goto _again
	tr10: cs = 13; goto _again
	tr9: cs = 14; goto _again

f1:
//line match.rl:19
 pb = p 
	goto _again
f0:
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
//line match.go:204
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
