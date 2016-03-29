# object-storage-cli
A Command Line (CLI) Tool for Utilizing Multiple Object Storage Systems from a Single Interface.

This CLI accepts a single flag called `--storage-type`. The value of that flag determines which object storage API to use, and the remaining sections in this document describe how the CLI gets the location, credentials and other information it needs to work with the specified system.

## `s3`

If the storage type is `s3`, the CLI reads three files which specify access key, access secret and region. Each file location can be configured by an environment variable. Each environment variable and its default is listed below.

- `ACCESS_KEY_FILE` (`/var/run/secrets/deis/objectstore/creds/accesskey`)
- `SECRET_KEY_FILE` (`/var/run/secrets/deis/objectstore/creds/secretkey`)
- `REGION_FILE` (`/var/run/secrets/deis/objectstore/creds/region`)
- `BUCKET_FILE` (`/var/run/secrets/deis/objectstore/creds/bucket`)

## `gcs`

If the storage type is `gcs`, the CLI reads two files which specify the bucket and GCS access key. Each file location can be configured by an environment variable. Each environment variable and its default is listed below.

- `KEY_FILE` (`/var/run/secrets/deis/objectstore/creds/key.json`)
	- This file should be a JSON encoded object and contain a `project-id` key, which specifies the GCS project ID
- `BUCKET_FILE` (`/var/run/secrets/deis/objectstore/creds/bucket`)

## `azure`

If the storage type is `azure`, the CLI reads three files which specify the account name, account key and container. Each file location can be configured by an environment variable. Each environment variable and its default is listed below.

- `ACCOUNT_NAME_FILE` (`/var/run/secrets/deis/objectstore/creds/accountname`)
- `ACCOUNT_KEY_FILE` (`/var/run/secrets/deis/objectstore/creds/accountkey`)
- `CONTAINER_FILE` (`/var/run/secrets/deis/objectstore/creds/container`)

## Empty String

If the storage type is not given, or is the empty string, the CLI assumes that it should use the AWS S3 API to talk to an object storage system.

In this case, the CLI requires information on where the S3 API compatible server is located, along with the expected authentication/authorization information.

It gets the location information from environment variables, and assumes that any value that starts with `$` is itself an environment variable. It also gets the auth information from three files whose locations are specified by environment variables as well. See below for the list of environment variables and their defaults.

- `S3_HOST` (`$DEIS_MINIO_SERVICE_HOST`)
- `S3_PORT` (`$DEIS_MINIO_SERVICE_PORT`)
- `ACCESS_KEY_FILE` (`/var/run/secrets/deis/objectstore/creds/accesskey`)
- `ACCESS_SECRET_FILE` (`/var/run/secrets/deis/objectstore/creds/secretkey`)
- `BUCKET_FILE` (`/var/run/secrets/deis/objectstore/creds/bucket`)
