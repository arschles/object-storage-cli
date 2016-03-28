package config

import (
	driver "github.com/deis/distribution/registry/storage/driver"
)

type Config interface {
	CreateDriver() (driver.StorageDriver, error)
}
