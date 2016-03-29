SHORT_NAME := object-storage-cli
REPO_PATH := github.com/arschles/${SHORT_NAME}
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.9.1
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}
BINARY_NAME := objstorage
TAGS := 'include_gcs'

bootstrap:
	${DEV_ENV_CMD} glide install

build:
	${DEV_ENV_CMD} go build -tags=${TAGS} -o ${BINARY_NAME}

build-mac:
	${DEV_ENV_PREFIX} -e GOOS=darwin -e GOARCH=amd64 ${DEV_ENV_IMAGE} go build -tags=${TAGS} -o ${BINARY_NAME}

test:
	${DEV_ENV_CMD} go test -tags=${TAGS} $$(glide nv)
