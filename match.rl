package gotenv

import (
	"fmt"
)

func match(data []byte) (e envvar, err error) {
	%% machine envvar;
	%% write data;

	cs, p, pe := 0, 0, len(data)
	eof := pe
	pb := 0

	text := func() string { return string(data[pb:p]) }
	textQuoted := func() string { return string(data[pb+1:p-1]) }

	%%{
		action mark { pb = p }

		action setKey { e.key = text() }
		action setVal { e.val = text() }
		action setValQuoted { e.val = textQuoted() }

		key = (alpha | '_')+ >mark (alnum | '_')* %setKey;
		nonQuote = (any - ('"'|'\''));
		valSingleQuoted  = '\'' >mark !'\'' '\'' %setValQuoted;
		valDoubleQuoted  = '"'  >mark !'"'  '"'  %setValQuoted;
		valSimple  = nonQuote >mark (any* nonQuote)? %setVal;
		val = (valSingleQuoted | valDoubleQuoted | valSimple | zlen);
		main := key '=' val;

		write init;
		write exec;
	}%%

	if cs < envvar_first_final {
		return e, fmt.Errorf("invalid state: %d", cs)
	}
	return e, nil
}
