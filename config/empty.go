package config

import (
	"github.com/deis/distribution/registry/storage/driver"
)

type Empty struct {
	AccessKeyFile    string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	AccessSecretFile string `envconfig:"ACCESS_SECRET_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
}

func (e Empty) CreateDriver() (driver.StorageDriver, error) {
	return nil, nil
}
