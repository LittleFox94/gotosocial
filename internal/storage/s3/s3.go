package s3

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/url"
	"time"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/superseriousbusiness/gotosocial/internal/config"
)

type s3Storage struct {
	client *minio.Client
	bucket string
}

func New(c *config.StorageConfig) (*s3Storage, error) {
	bucketLookupType := minio.BucketLookupDNS

	if c.S3.UsePathStyle {
		bucketLookupType = minio.BucketLookupPath
	}

	client, err := minio.New(c.S3.Endpoint, &minio.Options{
		Creds:        credentials.NewStaticV4(c.S3.AccessKeyID, c.S3.SecretAccessKey, ""),
		Secure:       c.S3.UseSSL,
		Region:       c.S3.Region,
		BucketLookup: bucketLookupType,
	})
	if err != nil {
		return nil, err
	}

	return &s3Storage{
		client: client,
		bucket: c.S3.Bucket,
	}, nil
}

func (s *s3Storage) Get(name string) ([]byte, error) {
	o, err := s.client.GetObject(context.TODO(), s.bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(o)
}

func (s *s3Storage) Put(name string, data []byte) error {
	_, err := s.client.PutObject(context.TODO(), s.bucket, name, bytes.NewBuffer(data), int64(len(data)), minio.PutObjectOptions{})
	return err
}

func (s *s3Storage) Delete(name string) error {
	return s.client.RemoveObject(context.TODO(), s.bucket, name, minio.RemoveObjectOptions{})
}

func (s *s3Storage) URL(name string) (*url.URL, error) {
	return s.client.PresignedGetObject(context.TODO(), s.bucket, name, 120*time.Second, url.Values{})
}
