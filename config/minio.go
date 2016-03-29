package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/docker/distribution/registry/storage/driver"
	s3 "github.com/docker/distribution/registry/storage/driver/s3-aws"
)

const (
	dollar      = '$'
	minioRegion = "us-east-1a"
)

var (
	errMissingHost = errors.New("missing s3 host")
	errMissingPort = errors.New("missing s3 port")
)

type Minio struct {
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
	accessKey, secretKey, bucket := files[0], files[1], files[2]
	params := s3.DriverParameters{
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		Bucket:         bucket,
		RegionEndpoint: fmt.Sprintf("%s:%s", host, port),
		Region:         minioRegion,
	}
	return s3.New(params)
}

// String is the fmt.Stringer interface implementation
func (e Minio) String() string {
	return MinioStorageType.String()
}
