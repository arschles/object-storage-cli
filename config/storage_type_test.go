package config

import (
	"testing"

	"github.com/arschles/assert"
)

func TestStorageTypeString(t *testing.T) {
	str := "abc"
	st := StorageType(str)
	assert.Equal(t, st.String(), str, "string value")
}

func TestStorageTypeFromString(t *testing.T) {
	type testCase struct {
		str      string
		expected StorageType
		err      bool
	}
	testCases := []testCase{
		testCase{str: "s3", expected: S3StorageType, err: false},
		testCase{str: "gcs", expected: GCSStorageType, err: false},
		testCase{str: "azure", expected: AzureStorageType, err: false},
		testCase{str: "minio", expected: AzureStorageType, err: false},
		testCase{str: "unknown", expected: "", err: true},
	}
	for _, tc := range testCases {
		st, err := StorageTypeFromString(tc.str)
		if tc.err && err == nil {
			t.Errorf("Expected error for %s, got none", tc.str)
			continue
		}
		if !tc.err && err != nil {
			t.Errorf("Expected no error for %s, got %s", tc.str, err)
			continue
		}
		if tc.err {
			assert.Equal(t, st.String(), "", "string value")
		} else {
			assert.Equal(t, st.String(), tc.str, "string value")
		}
	}
}
