package config

import (
	"github.com/docker/distribution/registry/storage/driver"
	gcs "github.com/docker/distribution/registry/storage/driver/gcs"
)

const (
	bucketParam  = "bucket"
	keyFileParam = "keyfile"
)

// GCS is the Config implementation for the GCS client
type GCS struct {
	KeyFile    string `envconfig:"KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/key.json"`
	BucketFile string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/bucket"`
}

// CreateDriver is the Config interface implementation
func (g GCS) CreateDriver() (driver.StorageDriver, error) {
	files, err := readFiles(true, g.BucketFile)
	if err != nil {
		return nil, err
	}
	bucket := files[0]
	params := map[string]interface{}{
		bucketParam:  bucket,
		keyFileParam: g.KeyFile,
	}
	return gcs.FromParameters(params)
}

// String is the fmt.Stringer interface implementation
func (g GCS) String() string {
	return GCSStorageType.String()
}
