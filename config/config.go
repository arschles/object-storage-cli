package config

import (
	"fmt"

	driver "github.com/docker/distribution/registry/storage/driver"
	"github.com/kelseyhightower/envconfig"
)

const (
	appName = "objstorage"
)

// Config is the generic interface from which a storage driver can be created
type Config interface {
	fmt.Stringer
	CreateDriver() (driver.StorageDriver, error)
}

type ErrUnsupportedStorageType struct {
	st StorageType
}

func (e ErrUnsupportedStorageType) Error() string {
	return fmt.Sprintf("Unsupported storage type %s", e.st)
}

// FromStorageType gets a Config implementation from the st param. If there's no known config for the given type, return ErrUnsupportedStorageType. Returns an appropriate error in all other error cases.
func FromStorageType(st StorageType) (Config, error) {
	var conf Config
	switch st {
	case S3StorageType:
		conf = new(S3)
	case GCSStorageType:
		conf = new(GCS)
	case AzureStorageType:
		conf = new(Azure)
	case EmptyStorageType:
		conf = new(Empty)
	default:
		return nil, ErrUnsupportedStorageType{st: st}
	}

	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// FromStorageTypeString calls StorageTypeFromString, then calls FromStorageType to get the Config from that storage type. Returns an appropriate error on all failures
func FromStorageTypeString(str string) (Config, error) {
	st, err := StorageTypeFromString(str)
	if err != nil {
		return nil, err
	}
	return FromStorageType(st)
}
