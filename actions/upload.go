package actions

import (
	"io/ioutil"
	"log"

	"github.com/arschles/object-storage-cli/config"
	"github.com/codegangsta/cli"
	"github.com/docker/distribution/context"
)

var (
	// UploadCommand is the cli.Command for use in the top level app commands list
	UploadCommand = cli.Command{
		Name:      "upload",
		ShortName: "up",
		Action:    Upload,
	}
)

// Upload is the cli handler for "upload" command
func Upload(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		log.Fatalf("This command should be called as 'upload $LOCAL_PATH $REMOTE_PATH'")
	}
	local := args[0]
	remote := args[1]

	conf, err := config.FromStorageTypeString(c.GlobalString(config.StorageTypeFlag))
	if err != nil {
		log.Fatal(err)
	}
	driver, err := conf.CreateDriver()
	if err != nil {
		log.Fatal(err)
	}

	fBytes, err := ioutil.ReadFile(local)
	if err != nil {
		log.Fatalf("Error reading local file %s (%s)", local, err)
	}
	ctx := context.Background()
	if err := driver.PutContent(ctx, remote, fBytes); err != nil {
		log.Fatalf("Error writing remote object %s (%s)", remote, err)
	}
	log.Printf("Successfully copied %s to %s", remote, local)
}
