package datasources

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	Client *s3.S3
}

func NewS3ClientWithCredentials(accessKey string, secretKey string, region string, roleArn string) (*S3Client, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}))

	var creds *credentials.Credentials
	if roleArn != "" {
		creds = stscreds.NewCredentials(sess, roleArn)
	} else {
		creds = credentials.NewStaticCredentials(accessKey, secretKey, "")
	}

	s3Client := s3.New(sess, &aws.Config{Credentials: creds})

	return &S3Client{Client: s3Client}, nil
}

func (c *S3Client) UploadFileToS3(bucket, key string, timeout time.Duration, fileContent interface{}) error {
	var body io.ReadSeeker

	// Determine byte slice or file object
	switch v := fileContent.(type) {
	case []byte:
		// Create a reader for the byte slice
		body = bytes.NewReader(v)
	case *os.File:
		body = v // File objects already implement io.Reader and io.ReadSeeker
	default:
		return fmt.Errorf("invalid file content type, must be []byte or *os.File")
	}

	// Create a context with timeout if specified
	ctx := context.Background()
	if timeout > 0 {
		var cancelFn func()
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
		defer cancelFn()
	}

	// Upload the file to S3
	_, err := c.Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		return fmt.Errorf("failed to upload object: %v", err)
	}

	fmt.Printf("Successfully uploaded file to %s/%s\n", bucket, key)
	return nil
}

func (c *S3Client) GetFileFromS3(bucket, key string) (*s3.GetObjectOutput, error) {
	output, err := c.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}

	fmt.Printf("Successfully retrieved file from %s/%s\n", bucket, key)
	return output, nil
}
