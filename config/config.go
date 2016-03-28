package config

import (
	driver "github.com/docker/distribution/registry/storage/driver"
)

type Config interface {
	CreateDriver() (driver.StorageDriver, error)
}
