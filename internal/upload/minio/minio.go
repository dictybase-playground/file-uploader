package minio

import (
	"fmt"
	"io/ioutil"

	upload "github.com/dictybase-playground/snapshot-upload/internal/upload"
	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
)

// Minio contains all configuration necessary to upload.
type Minio struct {
	bucket string
	folder string
	client *minio.Client
	logger *logrus.Entry
}

// NewUploadCreator acts as a constructor for Minio file uploads
func NewUploadCreator(bucket, folder string, client *minio.Client, logger *logrus.Entry) upload.Uploader {
	return &Minio{bucket: bucket, folder: folder, client: client, logger: logger}
}

// UploadFolder uploads files inside a folder to Minio
func (m *Minio) UploadFolder() error {
	found, err := m.client.BucketExists(m.bucket)
	if err != nil {
		m.logger.Errorf("could not find bucket %s", err)
		return fmt.Errorf("could not find bucket %s", err)
	}
	if found {
		m.logger.Info("Bucket found")
	}

	files, err := ioutil.ReadDir(m.folder)
	if err != nil {
		m.logger.Errorf("couldn't read directory %s", err)
		return fmt.Errorf("couldn't read directory %s", err)
	}

	for _, file := range files {
		filePath := m.folder + "/" + file.Name()
		n, err := m.client.FPutObject(m.bucket, file.Name(), filePath, minio.PutObjectOptions{})
		if err != nil {
			m.logger.Errorf("couldn't upload file %s", err)
			return fmt.Errorf("couldn't upload file %s", err)
		}
		m.logger.Infof("Successfully uploaded %s of size %d\n", file.Name(), n)
	}
	return nil
}
