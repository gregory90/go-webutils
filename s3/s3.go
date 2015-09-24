package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Upload(region string, bucket string, path string, file io.ReadSeeker) error {
	config := &aws.Config{
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
	}
	svc := s3.New(config)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
		Body:   file,
	}
	_, err := svc.PutObject(params)

	return err
}

func Get(region string, bucket string, path string) (io.ReadCloser, error) {
	config := &aws.Config{
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
	}
	svc := s3.New(config)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}
	resp, err := svc.GetObject(params)

	return resp.Body, err
}

func Delete(region string, bucket string, path string) error {
	config := &aws.Config{
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
	}
	svc := s3.New(config)

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}
	_, err := svc.DeleteObject(params)

	return err
}
