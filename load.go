package gotenv

import (
	"fmt"
	"os"
)

type envvar struct {
	key, val string
}

// Load given dotenv files. As a result, the env variables within
// the current process are set according to the assignments files;
// they should have a form `NAME=VALUE`.
// Whitespaces around equal sign `=`
// are allowed and ignored.
// If the value is surrounded by the double quotes `"`,
// it will be passed to the env unquoted.
// No value after equal sign means that the variable will be reset to empty.
// Hash sign `#` acts as a comment, ignoring everything until the end of line,
// unless it's within double quotes.
//
// Important: if the variable existed in the env prior to this call,
// it will *not* be set.
//
// If there is no given file, `.env` is assumed.
func Load(filenames ...string) (err error) {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	for _, fn := range filenames {
		f, err := os.Open(fn)
		if err != nil {
			return err
		}
		defer f.Close()

		ee, err := readenv(f)
		if err != nil {
			return fmt.Errorf("file %s: %w", fn, err)
		}
		ee = dedup(ee)

		for _, e := range ee {
			_, exists := os.LookupEnv(e.key)
			if exists {
				continue
			}
			err = os.Setenv(e.key, e.val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
