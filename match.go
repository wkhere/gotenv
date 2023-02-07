
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
	0, 0, 0, 0, 0, 5, 7, 0, 
	0, 7, 7, 0, 0, 7, 
}

const envvar_start int = 8
const envvar_first_final int = 8
const envvar_error int = 0

const envvar_en_main int = 8


//line match.rl:10

	cs, p, pe := 0, 0, len(data)
	eof := pe
	pb := 0

	text := func() string { return string(data[pb:p]) }
	textQuoted := func() string { return string(data[pb+1:p-1]) }

	
//line match.go:38
	{
	cs = envvar_start
	}

//line match.go:43
	{
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 8:
		switch data[p] {
		case 32:
			goto tr14;
		case 35:
			goto tr15;
		case 95:
			goto tr16;
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr14;
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr16;
			}
		default:
			goto tr16;
		}
		goto tr1;
	case 0:
		goto _out
	case 9:
		switch data[p] {
		case 32:
			goto tr17;
		case 35:
			goto tr18;
		case 95:
			goto tr16;
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr17;
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr16;
			}
		default:
			goto tr16;
		}
		goto tr1;
	case 10:
		switch data[p] {
		case 32:
			goto tr17;
		case 35:
			goto tr18;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr17;
		}
		goto tr1;
	case 11:
		goto tr1;
	case 1:
		switch data[p] {
		case 32:
			goto tr0;
		case 61:
			goto tr3;
		case 95:
			goto tr2;
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr0;
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr2;
				}
			case data[p] >= 65:
				goto tr2;
			}
		default:
			goto tr2;
		}
		goto tr1;
	case 2:
		switch data[p] {
		case 32:
			goto tr4;
		case 61:
			goto tr5;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr4;
		}
		goto tr1;
	case 12:
		switch data[p] {
		case 32:
			goto tr5;
		case 34:
			goto tr20;
		case 35:
			goto tr21;
		case 39:
			goto tr22;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr5;
		}
		goto tr19;
	case 13:
		switch data[p] {
		case 32:
			goto tr23;
		case 34:
			goto tr7;
		case 35:
			goto tr24;
		case 39:
			goto tr7;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr23;
		}
		goto tr6;
	case 3:
		switch data[p] {
		case 32:
			goto tr1;
		case 34:
			goto tr7;
		case 39:
			goto tr7;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1;
		}
		goto tr6;
	case 4:
		if data[p] == 34 {
			goto tr9;
		}
		goto tr8;
	case 5:
		if data[p] == 34 {
			goto tr10;
		}
		goto tr8;
	case 14:
		switch data[p] {
		case 32:
			goto tr25;
		case 34:
			goto tr10;
		case 35:
			goto tr26;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr25;
		}
		goto tr8;
	case 15:
		switch data[p] {
		case 32:
			goto tr27;
		case 34:
			goto tr10;
		case 35:
			goto tr28;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr27;
		}
		goto tr8;
	case 16:
		if data[p] == 34 {
			goto tr10;
		}
		goto tr8;
	case 17:
		switch data[p] {
		case 32:
			goto tr25;
		case 35:
			goto tr26;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr25;
		}
		goto tr8;
	case 6:
		if data[p] == 39 {
			goto tr12;
		}
		goto tr11;
	case 7:
		if data[p] == 39 {
			goto tr13;
		}
		goto tr11;
	case 18:
		switch data[p] {
		case 32:
			goto tr29;
		case 35:
			goto tr30;
		case 39:
			goto tr13;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr29;
		}
		goto tr11;
	case 19:
		switch data[p] {
		case 32:
			goto tr31;
		case 35:
			goto tr32;
		case 39:
			goto tr13;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr31;
		}
		goto tr11;
	case 20:
		if data[p] == 39 {
			goto tr13;
		}
		goto tr11;
	case 21:
		switch data[p] {
		case 32:
			goto tr29;
		case 35:
			goto tr30;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr29;
		}
		goto tr11;
	}

	tr1: cs = 0; goto _again
	tr2: cs = 1; goto _again
	tr16: cs = 1; goto f2
	tr4: cs = 2; goto _again
	tr0: cs = 2; goto f0
	tr7: cs = 3; goto _again
	tr20: cs = 4; goto f2
	tr8: cs = 5; goto _again
	tr22: cs = 6; goto f2
	tr11: cs = 7; goto _again
	tr14: cs = 8; goto _again
	tr15: cs = 9; goto f1
	tr17: cs = 10; goto _again
	tr23: cs = 10; goto f4
	tr18: cs = 11; goto f1
	tr5: cs = 12; goto _again
	tr3: cs = 12; goto f0
	tr6: cs = 13; goto _again
	tr19: cs = 13; goto f2
	tr21: cs = 13; goto f3
	tr24: cs = 13; goto f5
	tr10: cs = 14; goto _again
	tr27: cs = 15; goto _again
	tr25: cs = 15; goto f6
	tr28: cs = 16; goto f1
	tr26: cs = 16; goto f7
	tr9: cs = 17; goto _again
	tr13: cs = 18; goto _again
	tr31: cs = 19; goto _again
	tr29: cs = 19; goto f6
	tr32: cs = 20; goto f1
	tr30: cs = 20; goto f7
	tr12: cs = 21; goto _again

f2:
//line match.rl:19
 pb = p 
	goto _again
f0:
//line match.rl:21
 e.key = text() 
	goto _again
f4:
//line match.rl:22
 e.val = text() 
	goto _again
f6:
//line match.rl:23
 e.val = textQuoted() 
	goto _again
f1:
//line match.rl:25
 return e, nil 
	goto _again
f3:
//line match.rl:19
 pb = p 
//line match.rl:25
 return e, nil 
	goto _again
f5:
//line match.rl:22
 e.val = text() 
//line match.rl:25
 return e, nil 
	goto _again
f7:
//line match.rl:23
 e.val = textQuoted() 
//line match.rl:25
 return e, nil 
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
		case 5:
//line match.rl:22
 e.val = text() 
		case 7:
//line match.rl:23
 e.val = textQuoted() 
//line match.go:388
		}
	}

	_out: {}
	}

//line match.rl:39


	if cs < envvar_first_final {
		return e, fmt.Errorf("invalid state: %d", cs)
	}
	return e, nil
}
