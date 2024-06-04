package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Cloud interface {
	Upload(ctx context.Context, file io.Reader, dir string) (*s3.PutObjectOutput, string, error)
	DownloadLargeObject(objectKey string) ([]byte, error)
	UploadLargeObject(objectKey string, largeObject []byte) (string, error)
}
