package upload

// Uploader is an interface for uploading files
type Uploader interface {
	// UploadSnapshots uploads to online storage
	UploadSnapshots() error
}
