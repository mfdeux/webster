package store

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

type FileStore interface {
	GetFile() ([]byte, error)
	GetFileURL() (string, error)
	PutFile() error
	DeleteFile() error
}

type DiskFileStore struct {
	rootPath string
}

func NewDiskFileStore(rootPath string) (*DiskFileStore, error) {
	return &DiskFileStore{
		rootPath: rootPath,
	}, nil
}

func (s *DiskFileStore) GetFile(key string) ([]byte, error) {
	return ioutil.ReadFile(key)
}

func (s *DiskFileStore) GetFileURL(key string) (string, error) {
	return "", nil
}

func (s *DiskFileStore) DeleteFile(key string) error {
	return os.Remove(key)
}

func (s *DiskFileStore) PutFile(key string, file io.Reader) error {
	return ioutil.WriteFile(key, []byte("write this to file"), 0644)
}

type AWSFileStore struct {
	config     aws.Config
	client     *s3.Client
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	bucket     string
}

func NewAWSFileStore(accessKey, secretKey, bucket string) (*AWSFileStore, error) {
	creds := aws.NewStaticCredentialsProvider(accessKey, secretKey, "default")
	cfg := aws.Config{
		Credentials: creds,
		Region:      endpoints.UsEast1RegionID,
	}

	svc := s3.New(cfg)
	uploader := s3manager.NewUploader(cfg)
	downloader := s3manager.NewDownloader(cfg)

	return &AWSFileStore{
		config:     cfg,
		client:     svc,
		uploader:   uploader,
		downloader: downloader,
		bucket:     bucket,
	}, nil
}

func (s *AWSFileStore) GetFile(key string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := s.downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *AWSFileStore) GetFileURL(key string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}
	req := s.client.GetObjectRequest(input)
	return req.Presign(15 * time.Minute)
}

func (s *AWSFileStore) DeleteFile(key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}
	req := s.client.DeleteObjectRequest(input)
	_, err := req.Send(context.Background())
	return err
}

func (s *AWSFileStore) PutFile(key string, file io.Reader) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    s3.ObjectCannedACLPrivate,
	})
	return err
}
