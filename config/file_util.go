package config

import (
	"io/ioutil"
	"strings"
)

func readFiles(trimSpaces bool, names ...string) ([]string, error) {
	ret := make([]string, len(names))
	for i, name := range names {
		b, err := ioutil.ReadFile(name)
		if err != nil {
			return nil, err
		}
		if trimSpaces {
			ret[i] = strings.TrimSpace(string(b))
		} else {
			ret[i] = string(b)
		}
	}

	return ret, nil
}
