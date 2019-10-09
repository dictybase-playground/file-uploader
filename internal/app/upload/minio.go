package upload

import (
	log "github.com/dictybase-playground/file-uploader/internal/logger"
	m "github.com/dictybase-playground/file-uploader/internal/upload/minio"
	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
)

func UploadFilesMinio(c *cli.Context) error {
	useSSL := true
	logger, err := log.GetLogger(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	// Initialize minio client object
	minioClient, err := minio.New(c.String("minio-endpoint"), c.String("minio-access-key"), c.String("minio-secret-key"), useSSL)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	m.NewUploadCreator(c.String("minio-bucket"), c.String("folder"), minioClient, logger)
	return nil
}
