package main

import (
	"os"

	"github.com/arschles/object-storage-cli/actions"
	"github.com/arschles/object-storage-cli/config"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "objstorage"
	app.Usage = "Use a variety of different object storage systems with a single tool"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  config.StorageTypeFlag,
			Value: config.S3StorageType.String(),
			Usage: "Specify the type of the object storage system",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:      "download",
			ShortName: "dl",
			Action:    actions.Download,
		},
	}
	app.Run(os.Args)
}
