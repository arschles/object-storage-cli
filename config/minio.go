package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/docker/distribution/registry/storage/driver"
	"github.com/docker/distribution/registry/storage/driver/factory"
	// this blank import is used to register the S3 driver with the storage driver factory
	_ "github.com/docker/distribution/registry/storage/driver/s3-aws"
)

const (
	dollar = '$'
)

var (
	errMissingHost = errors.New("missing s3 host")
	errMissingPort = errors.New("missing s3 port")
)

// Minio is the Config implementation for the Minio client
type Minio struct {
	AccessKeyFile    string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	AccessSecretFile string `envconfig:"ACCESS_SECRET_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
	BucketFile       string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/bucket"`
	S3Host           string `envconfig:"S3_HOST" default:"$DEIS_MINIO_SERVICE_HOST"`
	S3Port           string `envconfig:"S3_PORT" default:"$DEIS_MINIO_SERVICE_PORT"`
	Region           string `envconfig:"REGION" default:"us-east-1"`
	Secure           bool   `envconfig:"SECURE" default:"false"`
	V4Auth           bool   `envconfig:"V4_AUTH" default:"true"`
}

func parseEnvVar(val string) string {
	if val[0] == dollar {
		return os.Getenv(val[1:])
	}
	return val
}

// CreateDriver is the Config interface implementation
func (e Minio) CreateDriver() (driver.StorageDriver, error) {
	files, err := readFiles(true, e.AccessKeyFile, e.AccessSecretFile, e.BucketFile)
	if err != nil {
		return nil, err
	}
	host := parseEnvVar(e.S3Host)
	if host == "" {
		return nil, errMissingHost
	}
	port := parseEnvVar(e.S3Port)
	if port == "" {
		return nil, errMissingPort
	}
	key, secret, bucket := files[0], files[1], files[2]
	params := map[string]interface{}{
		"accesskey":      key,
		"secretkey":      secret,
		"region":         e.Region,
		"bucket":         bucket,
		"regionendpoint": fmt.Sprintf("http://%s:%s", host, port),
		"secure":         e.Secure,
		"v4auth":         e.V4Auth,
	}
	return factory.Create("s3aws", params)
}

// String is the fmt.Stringer interface implementation
func (e Minio) String() string {
	return MinioStorageType.String()
}
