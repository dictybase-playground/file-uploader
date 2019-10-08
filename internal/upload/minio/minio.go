package minio

import (
	upload "github.com/dictybase-playground/snapshot-upload/internal/upload"

	"github.com/sirupsen/logrus"
)

// Minio contains all configuration necessary to upload.
type Minio struct {
	endpoint  string
	accessKey string
	secretKey string
	bucket    string
	folder    string
	logger    *logrus.Entry
}

// NewUploadCreator acts as a constructor for Minio file uploads
func NewUploadCreator(endpoint, accessKey, secretKey, bucket, folder string, logger *logrus.Entry) upload.Uploader {
	return &Minio{endpoint: endpoint, accessKey: accessKey, secretKey: secretKey, bucket: bucket, folder: folder, logger: logger}
}

// UploadSnapshots uploads snapshots to Minio
func (m *Minio) UploadSnapshots() error {

	return nil
}
