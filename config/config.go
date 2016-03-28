package config

import (
	driver "github.com/docker/distribution/registry/storage/driver"
	"github.com/kelseyhightower/envconfig"
)

// Config is the generic interface from which a storage driver can be created
type Config interface {
	CreateDriver() (driver.StorageDriver, error)
}

// GetConfigFromStorageType gets a Config implementation from the st param. If there's no known config
// for the given type, returns an appropriate error
func GetConfigFromStorageType(st StorageType) (Config, error) {
	// NYI
	return nil, nil
}
