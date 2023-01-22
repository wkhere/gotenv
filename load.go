package gotenv

import (
	"fmt"
	"os"
)

type envvar struct {
	key, val string
}

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
