package upload

// Uploader is an interface for uploading files
type Uploader interface {
	// UploadFolder uploads files inside a folder to Minio
	UploadFolder() error
}
