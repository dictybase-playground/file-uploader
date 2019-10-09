package main

import (
	"os"

	"github.com/dictybase-playground/snapshot-upload/internal/app/upload"
	"github.com/dictybase-playground/snapshot-upload/internal/app/validate"
	cli "github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "file-uploader"
	app.Usage = "cli for uploading files to online storage"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-format",
			Usage: "format of the logging out, either of json or text.",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "log level for the application",
			Value: "debug",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "minio",
			Usage:  "uploads files to minio s3 storage",
			Action: upload.UploadFilesMinio,
			Before: validate.ValidateMinioArgs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "minio-endpoint",
					Usage:  "minio endpoint",
					Value:  "ericstorage.dictybase.dev",
					EnvVar: "MINIO_ENDPOINT",
				},
				cli.StringFlag{
					Name:   "minio-access-key",
					Usage:  "minio access key",
					EnvVar: "MINIO_ACCESS_KEY",
				},
				cli.StringFlag{
					Name:   "minio-secret-key",
					Usage:  "minio secret key",
					EnvVar: "MINIO_SECRET_KEY",
				},
				cli.StringFlag{
					Name:  "minio-bucket",
					Usage: "minio bucket to upload to",
				},
				cli.StringFlag{
					Name:  "folder, f",
					Usage: "folder to read and upload files from",
				},
			},
		},
	}
	app.Run(os.Args)
}
