package config

import (
	"fmt"

	"github.com/docker/distribution/registry/storage/driver"
	s3 "github.com/docker/distribution/registry/storage/driver/s3-aws"
)

type S3 struct {
	AccessKeyFile string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	SecretKeyFile string `envconfig:"SECRET_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
	RegionFile    string `envconfig:"REGION_FILE" default:"/var/run/secrets/deis/objectstore/creds/region"`
	BucketFile    string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/bucket"`
}

func endpointURL(region string) string {
	return fmt.Sprintf("s3-%s.amazonaws.com", region)
}

// CreateDriver is the Config interface implementation
func (s S3) CreateDriver() (driver.StorageDriver, error) {
	files, err := readFiles(true, s.AccessKeyFile, s.SecretKeyFile, s.RegionFile, s.BucketFile)
	if err != nil {
		return nil, err
	}
	keyBytes, secretBytes, regionBytes, bucketBytes := files[0], files[1], files[2], files[3]
	params := s3.DriverParameters{
		AccessKey:      string(keyBytes),
		SecretKey:      string(secretBytes),
		Bucket:         string(bucketBytes),
		Region:         string(regionBytes),
		RegionEndpoint: endpointURL(string(regionBytes)),
	}
	return s3.New(params)
}

// String is the fmt.Stringer interface implementation
func (s S3) String() string {
	return S3StorageType.String()
}
