package upload

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/minio/minio-go/v6"
	"github.com/urfave/cli"
)

func UploadSnapshotsMinio(c *cli.Context) error {
	endpoint := c.String("minio-endpoint")
	accessKeyID := c.String("minio-access-key")
	secretAccessKey := c.String("minio-secret-key")
	bucket := c.String("minio-bucket")
	folder := c.String("folder")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	found, err := minioClient.BucketExists(bucket)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if found {
		fmt.Println("Bucket found")
	}

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		n, err := minioClient.FPutObject(bucket, file.Name(), file.Name(), minio.PutObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Successfully uploaded %s of size %d\n", file.Name(), n)
	}

	return nil
}
