package main

import (
	"fmt"
)

type s3Config struct {
	AccessKeyFile string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	SecretKeyFile string `envconfig:"SECRET_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
	RegionFile    string `envconfig:"/var/run/secrets/deis/objectstore/creds/region" default:"/var/run/secrets/deis/objectstore/creds/region"`
}

type gcsConfig struct {
	KeyFile    string `envconfig:"KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/key.json"`
	BucketFile string `envconfig:"BUCKET_FILE" default:"/var/run/secrets/deis/objectstore/creds/builder-bucket"`
}

type azureConfig struct {
	AccountNameFile string `envconfig:"ACCOUNT_NAME_FILE" default:"/var/run/secrets/deis/objectstore/creds/accountname"`
	AccountKeyFile  string `envconfig:"ACCOUNT_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accountkey"`
	ContainerFile   string `envconfig:"CONTAINER_FILE" default:"/var/run/secrets/deis/objectstore/creds/container"`
}

type emptyConfig struct {
	AccessKeyFile    string `envconfig:"ACCESS_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accesskey"`
	AccessSecretFile string `envconfig:"ACCESS_SECRET_FILE" default:"/var/run/secrets/deis/objectstore/creds/secretkey"`
}
