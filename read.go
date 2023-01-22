package gotenv

import (
    "bufio"
    "fmt"
    "io"
)

func readenv(r io.Reader) (ee []envvar, _ error) {
    b := bufio.NewScanner(r)
    lineno := 0
    for b.Scan() {
        lineno++
        e, err := match(b.Bytes())
        if err != nil {
            return ee, fmt.Errorf(
                "no match at line %d: %v\n  parsed so far: %v",
                lineno, err, e,
            )
        }
        ee = append(ee, e)
    }
    return ee, b.Err()
}
