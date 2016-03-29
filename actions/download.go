package actions

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/arschles/object-storage-cli/config"
	"github.com/codegangsta/cli"
	"github.com/docker/distribution/context"
)

var (
	// DownloadCommand is the cli.Command for use in the top level app commands list
	DownloadCommand = cli.Command{
		Name:      "download",
		ShortName: "dl",
		Action:    Download,
	}
)

// Download is the cli handler for "download" command
func Download(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		log.Fatalf("This command should be called as 'download $REMOTE_PATH $LOCAL_PATH'")
	}

	remote := args[0]
	local := args[1]

	conf, err := config.FromStorageTypeString(c.GlobalString(config.StorageTypeFlag))
	if err != nil {
		log.Fatal(err)
	}
	driver, err := conf.CreateDriver()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	b, err := driver.GetContent(ctx, remote)
	if err != nil {
		log.Fatalf("Error downloading %s (%s)", remote, err)
	}
	if err := ioutil.WriteFile(local, b, os.ModePerm); err != nil {
		log.Fatalf("Error writing %s to %s (%s)", remote, local, err)
	}

	log.Printf("Successfully copied %s to %s", remote, local)
}
