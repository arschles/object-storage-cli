package actions

import (
	"github.com/arschles/object-storage-cli/config"
	"github.com/codegangsta/cli"
)

// Download is the cli handler for "download" command
func Download(c *cli.Context) {
	storageType := config.StorageTypeFromString(c.GlobalString(config.StorageTypeFlag))
}
