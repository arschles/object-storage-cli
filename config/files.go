package config

import (
	"io/ioutil"
)

func readFiles(names ...string) ([][]byte, error) {
	ret := make([][]byte, len(names))
	for i, name := range names {
		b, err := ioutil.ReadFile(name)
		if err != nil {
			return nil, err
		}
		ret[i] = b
	}

	return ret, nil
}
