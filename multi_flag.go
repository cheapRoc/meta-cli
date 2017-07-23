package main

import (
	"fmt"
	"strings"
)

type MultiFlag struct {
	Values map[string]string
}

func (f MultiFlag) String() string {
	return fmt.Sprintf("%v", f.Values)
}

func (f *MultiFlag) Set(value string) error {
	if f.Len() == 0 {
		f.Values = make(map[string]string, 1)
	}
	pair := strings.Split(value, "=")
	if len(pair) < 2 {
		return fmt.Errorf(
			"flag value '%v' was not in the format 'key=val'", value)
	}
	key, val := strings.Join(pair[0:1], ""), strings.Join(pair[1:2], "")
	f.Values[key] = val
	return nil
}

func (f MultiFlag) Len() int {
	return len(f.Values)
}
