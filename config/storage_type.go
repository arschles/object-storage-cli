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

// ErrUnknownStorageType is an error that indicates that a given storage type string is unknown
type ErrUnknownStorageType struct {
	typeStr string
}

// Error is the error interface implementation
func (e ErrUnknownStorageType) Error() string {
	return fmt.Sprintf("Unknown storage type %s", e.typeStr)
}

// StorageTypeFromString attempts to convert s into a StorageType value. Returns ErrUnknownStorageType if s doesn't correspond to any supported storage type
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
		return EmptyStorageType, ErrUnknownStorageType{typeStr: s}
	}
}
