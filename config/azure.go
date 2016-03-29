package config

import (
	"github.com/docker/distribution/registry/storage/driver"
	azure "github.com/docker/distribution/registry/storage/driver/azure"
)

type Azure struct {
	AccountNameFile string `envconfig:"ACCOUNT_NAME_FILE" default:"/var/run/secrets/deis/objectstore/creds/accountname"`
	AccountKeyFile  string `envconfig:"ACCOUNT_KEY_FILE" default:"/var/run/secrets/deis/objectstore/creds/accountkey"`
	ContainerFile   string `envconfig:"CONTAINER_FILE" default:"/var/run/secrets/deis/objectstore/creds/container"`
}

// CreateDriver is the Config interface implementation
func (a Azure) CreateDriver() (driver.StorageDriver, error) {
	files, err := readFiles(a.AccountNameFile, a.AccountKeyFile, a.ContainerFile)
	if err != nil {
		return nil, err
	}
	accountNameBytes, accountKeyBytes, containerBytes := files[0], files[1], files[2]
	return azure.New(string(accountNameBytes), string(accountKeyBytes), string(containerBytes), "")
}

// Name is the fmt.Stringer interface implementation
func (a Azure) String() string {
	return AzureStorageType.String()
}
