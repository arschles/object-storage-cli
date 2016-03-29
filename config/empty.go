package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/docker/distribution/registry/storage/driver"
	s3 "github.com/docker/distribution/registry/storage/driver/s3-aws"
)

const (
	dollar = '$'
)

var (
	errMissingHost = errors.New("missing s3 host")
	errMissingPort = errors.New("missing s3 port")
)

type Empty struct {
	AccessKeyFile    string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	AccessSecretFile string `envconfig:"ACCESS_SECRET_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
	BucketFile       string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/bucket"`
	S3Host           string `envconfig:"S3_HOST" default:"$DEIS_MINIO_SERVICE_HOST"`
	S3Port           string `envconfig:"S3_PORT" default:"$DEIS_MINIO_SERVICE_PORT"`
}

func parseEnvVar(val string) string {
	if val[0] == dollar {
		return os.Getenv(val[1:])
	}
	return val
}

// CreateDriver is the Config interface implementation
func (e Empty) CreateDriver() (driver.StorageDriver, error) {
	files, err := readFiles(e.AccessKeyFile, e.AccessSecretFile, e.BucketFile)
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
	params := s3.DriverParameters{
		AccessKey:      string(files[0]),
		SecretKey:      string(files[1]),
		Bucket:         string(files[2]),
		RegionEndpoint: fmt.Sprintf("%s:%s", host, port),
	}
	return s3.New(params)
}

// String is the fmt.Stringer interface implementation
func (e Empty) String() string {
	return EmptyStorageType.String()
}
