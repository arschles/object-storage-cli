package config

import (
	"io/ioutil"

	"github.com/docker/distribution/registry/storage/driver"
	gcs "github.com/docker/distribution/registry/storage/driver/gcs"
)

const (
	bucketParam  = "bucket"
	keyFileParam = "keyfile"
)

type GCS struct {
	KeyFile    string `envconfig:"KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/key.json"`
	BucketFile string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/builder-bucket"`
}

// CreateDriver is the Config interface implementation
func (g GCS) CreateDriver() (driver.StorageDriver, error) {
	bucketBytes, err := ioutil.ReadFile(g.BucketFile)
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params[bucketParam] = string(bucketBytes)
	params[keyFileParam] = g.KeyFile
	return gcs.FromParameters(params)
}
