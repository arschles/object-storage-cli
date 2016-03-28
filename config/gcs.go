package config

import (
	"github.com/deis/distribution/registry/storage/driver"
)

type GCS struct {
	KeyFile    string `envconfig:"KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/key.json"`
	BucketFile string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/builder-bucket"`
}

func (g GCS) CreateDriver() (driver.StorageDriver, error) {

}
