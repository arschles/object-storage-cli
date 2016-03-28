package config

import (
  "io/ioutil"

	"github.com/deis/distribution/registry/storage/driver"
	s3 "github.com/deis/distribution/registry/storage/driver/s3-aws"
)

type S3 struct {
	AccessKeyFile string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	SecretKeyFile string `envconfig:"SECRET_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
	RegionFile    string `envconfig:"REGION_FILE" default:"/var/run/secrets/deis/objectstore/creds/region"`
  BucketFile string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/builder-bucket"`
}

func (s S3) CreateDriver() (driver.StorageDriver, error) {
  keyBytes, err := ioutil.ReadFile(s.AccessKeyFile)
  if err != nil {
    return nil, err
  }
  secretBytes, err := ioutil.ReadFile(s.SecretKeyFile)
  if err != nil {
    return nil, err
  }
  regionBytes, err := ioutil.ReadFile(s.RegionFile)
  if err != nil {
    return nil, err
  }
  bucketBytes, err := ioutil.ReadFile(s.BucketFile)
  if err != nil {
    return nil, err
  }
  params := s3.DriverParameters{
    AccessKey: string(keyBytes),
    SecretKey: string(secretBytes),
    Bucket: string(bucketBytes),
    Region: string(regionBytes)
  }
}
