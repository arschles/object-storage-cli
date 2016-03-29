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
	// S3StorageType is the storage type for Amazon S3. See https://aws.amazon.com/s3/ for more information
	S3StorageType StorageType = "s3"
	// GCSStorageType is the storage type for Google Cloud Storage. See https://cloud.google.com/storage/ for more information
	GCSStorageType StorageType = "gcs"
	// AzureStorageType is the storage type for Azure Blob Storage. See https://azure.microsoft.com/en-us/services/storage/ for more information
	AzureStorageType StorageType = "azure"
	// MinioStorageType is the storage type for Minio. See https://minio.io/ for more information
	MinioStorageType StorageType = "minio"
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
	case MinioStorageType.String():
		return MinioStorageType, nil
	default:
		return "", ErrUnknownStorageType{typeStr: s}
	}
}
