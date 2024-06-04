package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3ConfigInfo struct {
	BucketName       string
	ShareEndPoint    string
	DefaultDirectory string
}

type S3Service struct {
	client       *s3.Client
	s3ConfigInfo S3ConfigInfo
}

func NewS3CloudService(client *s3.Client, cfg *S3ConfigInfo) *S3Service {
	return &S3Service{
		client:       client,
		s3ConfigInfo: *cfg,
	}
}

// Upload reads from a file and puts the data into an object in a bucket.
func (s *S3Service) Upload(ctx context.Context, file io.Reader, dir string) (*s3.PutObjectOutput, string, error) {
	fileBuff, object, contentType, err := s.uploadValidate(file)
	if err != nil {
		return nil, "", err
	}

	if dir != "" {
		object = fmt.Sprintf("%s/%s", dir, object) // append root dir in front of object
	}

	objectOutput, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.s3ConfigInfo.BucketName),
		Key:         aws.String(object),
		Body:        fileBuff,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return nil, "", err
	}

	return objectOutput, object, nil
}

func (s *S3Service) getFileContentType(f []byte) (string, error) {
	return http.DetectContentType(f), nil
}

func (s *S3Service) uploadValidate(file io.Reader) (*bytes.Buffer, string, string, error) {
	var (
		buff   = &bytes.Buffer{}
		object = uuid.New().String()
	)

	if _, err := buff.ReadFrom(file); err != nil {
		return nil, "", "", fmt.Errorf("failed to to read buff from file, error: %v", err)
	}

	contentType, err := s.getFileContentType(buff.Bytes())
	if err != nil {
		return nil, "", "", err
	}

	return buff, object, contentType, nil
}

// DownloadLargeObject uses a download manager to download an object from a bucket.
// The download manager gets the data in parts and writes them to a buffer until all of
// the data has been downloaded.
func (s *S3Service) DownloadLargeObject(objectKey string) ([]byte, error) {
	var partMiBs int64 = 50
	downloader := manager.NewDownloader(s.client, func(d *manager.Downloader) {
		d.PartSize = partMiBs * 1024 * 1024
	})
	buffer := manager.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(s.s3ConfigInfo.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't download large object from %v:%v. Here's why: %v\n",
			s.s3ConfigInfo.BucketName, objectKey, err)
	}
	return buffer.Bytes(), err
}

// UploadLargeObject uses an upload manager to upload data to an object in a bucket.
// The upload manager breaks large data into parts and uploads the parts concurrently.
func (s *S3Service) UploadLargeObject(objectKey string, largeObject []byte) (string, error) {
	largeBuffer := bytes.NewReader(largeObject)
	var partMiBs int64 = 10
	uploader := manager.NewUploader(s.client, func(u *manager.Uploader) {
		u.PartSize = partMiBs * 1024 * 1024
	})
	contentType, err := s.getFileContentType(largeObject)
	if err != nil {
		return "", err
	}
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.s3ConfigInfo.BucketName),
		Key:         aws.String(objectKey),
		Body:        largeBuffer,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Printf("Couldn't upload large object to %v:%v. Here's why: %v\n",
			s.s3ConfigInfo.BucketName, objectKey, err)
	}

	return objectKey, err
}
