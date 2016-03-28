package config

import (
	"github.com/deis/distribution/registry/storage/driver"
)

type Azure struct {
	AccountNameFile string `envconfig:"ACCOUNT_NAME_FILE" default:"/var/run/secrets/deis/objectstore/creds/accountname"`
	AccountKeyFile  string `envconfig:"ACCOUNT_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accountkey"`
	ContainerFile   string `envconfig:"CONTAINER_FILE" default:"/var/run/secrets/deis/objectstore/creds/container"`
}

func (a Azure) CreateDriver() (driver.StorageDriver, error) {
	return nil, nil
}
