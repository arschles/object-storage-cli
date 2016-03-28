package config

import (
	"fmt"
)

// StorageType represents the type of storage that a CLI command should work with
type StorageType string

func (c StorageType) String() string {
	return string(c)
}

const (
	S3StorageType    StorageType = "s3"
	GCSStorageType   StorageType = "gcs"
	AzureStorageType StorageType = "azure"
	EmptyStorageType StorageType = ""
)

type ErrUnsupportedStorageType struct {
	typeStr string
}

func (e ErrUnsupportedStorageType) Error() string {
	return fmt.Sprintf("Unsupported storage type %s", e.typeStr)
}

func StorageTypeFromString(s string) (StorageType, error) {
	switch s {
	case S3StorageType.String():
		return S3StorageType, nil
	case GCSStorageType.String():
		return GCSStorageType, nil
	case AzureStorageType.String():
		return AzureStorageType, nil
	case EmptyStorageType.String():
		return EmptyStorageType, nil
	default:
		return EmptyStorageType, ErrUnsupportedStorageType{typeStr: s}
	}
}
