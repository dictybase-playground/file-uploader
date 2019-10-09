package upload

import (
	"io/ioutil"

	log "github.com/dictybase-playground/file-uploader/internal/logger"
	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
)

func UploadFilesMinio(c *cli.Context) error {
	endpoint := c.String("minio-endpoint")
	accessKey := c.String("minio-access-key")
	secretKey := c.String("minio-secret-key")
	bucket := c.String("minio-bucket")
	folder := c.String("folder")
	useSSL := true

	logger, err := log.GetLogger(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	minioClient, err := minio.New(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	found, err := minioClient.BucketExists(bucket)
	if err != nil {
		logger.Errorf("could not find bucket %s", err)
		return cli.NewExitError(err.Error(), 2)
	}
	if found {
		logger.Info("Bucket found")
	}

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		logger.Errorf("couldn't read directory %s", err)
		return cli.NewExitError(err.Error(), 2)
	}

	for _, file := range files {
		filePath := folder + "/" + file.Name()
		n, err := minioClient.FPutObject(bucket, file.Name(), filePath, minio.PutObjectOptions{})
		if err != nil {
			logger.Errorf("couldn't upload file %s", err)
			return cli.NewExitError(err.Error(), 2)
		}
		logger.Infof("Successfully uploaded %s of size %d\n", file.Name(), n)
	}

	return nil
}
