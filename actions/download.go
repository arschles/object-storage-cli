package actions

import (
	"io"
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
		log.Fatalf("This command should be called as 'objstorage download $REMOTE_PATH $LOCAL_PATH'")
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

	fd, err := os.Create(local)
	if err != nil {
		log.Fatalf("Error creating/overwriting local file %s (%s)", local, err)
	}
	ctx := context.Background()
	rdr, err := driver.Reader(ctx, remote, 0)
	if err != nil {
		log.Fatalf("Error finding remote object %s (%s)", remote, err)
	}
	defer rdr.Close()
	if _, err := io.Copy(fd, rdr); err != nil {
		log.Fatalf("Error copying remote %s to local %s (%s)", remote, local, err)
	}
	log.Printf("Successfully copied %s to %s", remote, local)
}
