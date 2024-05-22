// Package gotenv implements reading dotenv files and setting env variables
// based on their content.
package gotenv

import (
	"fmt"
	"os"
)

type EnvVar struct {
	Key, Val string
}

// Read reads dotenv files.
//
// Files should contain variable definitions in a form `KEY=VALUE`.
// Whitespaces at the beginning of the line and around equal sign `=`
// are allowed and ignored.
// If the value is surrounded by the single `'` or double  `"` quotes,
// it will be passed to the env unquoted.
// No value after equal sign means that the variable will be reset to empty.
// Hash sign `#` acts as a comment, ignoring everything until the end of line,
// unless it's within quotes.
// If there are multiple definitions for the same KEY, in a single file or
// spanning multiple files, the last definition wins.
//
// If there is no given file, `.env` is assumed.
func Read(filenames ...string) (env []EnvVar, _ error) {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	for _, fn := range filenames {
		f, err := os.Open(fn)
		if err != nil {
			return env, err
		}
		defer f.Close()

		ee, err := readenv(f)
		if err != nil {
			return env, fmt.Errorf("file %s: %w", fn, err)
		}
		env = append(env, ee...)
	}

	return dedup(env), nil
}

// Setenv sets the env variables within the current process
// according to the given `env` argument, most likely produced by [Read].
//
// Important: if the variable existed in the process env prior to this call,
// it will *not* be set.
func Setenv(env []EnvVar) error {
	for _, e := range env {
		_, exists := os.LookupEnv(e.Key)
		if exists {
			continue
		}
		err := os.Setenv(e.Key, e.Val)
		if err != nil {
			return err
		}
	}
	return nil
}

// Load given dotenv files.
//
// Load combines [Read] and [Setenv].
// As a result, the env variables within the current process are set
// according to the definitions in given files;
// see [Read] for the notes on syntax.
func Load(filenames ...string) error {
	env, err := Read(filenames...)
	if err != nil {
		return err
	}

	for _, e := range env {
		_, exists := os.LookupEnv(e.Key)
		if exists {
			continue
		}
		err = os.Setenv(e.Key, e.Val)
		if err != nil {
			return err
		}
	}
	return nil
}
