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
	// UploadCommand is the cli.Command for use in the top level app commands list
	UploadCommand = cli.Command{
		Name:      "upload",
		ShortName: "up",
		Action:    Download,
	}
)

// Upload is the cli handler for "download" command
func Upload(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		log.Fatalf("This command should be called as 'objstorage upload $LOCAL_PATH $REMOTE_PATH'")
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

	fd, err := os.Open(local)
	if err != nil {
		log.Fatalf("Error opening local file %s for read (%s)", local, err)
	}
	ctx := context.Background()
	fw, err := driver.Writer(ctx, remote, false)
	if err != nil {
		if err := fw.Cancel(); err != nil {
			log.Printf("Cancelling the write operation failed while finding remote object %s (%s)", remote, err)
		}
		log.Fatalf("Error finding remote object %s (%s)", remote, err)
	}
	defer fw.Close()
	if _, err := io.Copy(fw, fd); err != nil {
		if err := fw.Cancel(); err != nil {
			log.Printf("Cancelling the write operation failed while writing %s to the remote object %s (%s)", local, remote, err)
		}
		log.Fatalf("Error copying local %s to remote %s (%s)", local, remote, err)
	}
	if err := fw.Commit(); err != nil {
		log.Fatalf("Error committing local %s to remote %s transfer operation (%s)", local, remote, err)
	}
	log.Printf("Successfully copied %s to %s", remote, local)
}
