// Package awssdk provides functions for AWS product integrations.
package awssdk

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"

	"gin-starter/common/helper"
	"gin-starter/config"
)

// S3Bucket wraps AWS S3 client operations.
type S3Bucket struct {
	cfg    config.Config
	client *s3.Client
}

// NewS3Bucket initializes an S3 client using the provided config.
// If region or credentials are missing in AWS config, it falls back
// to environment variables, shared config files, or IAM roles.
func NewS3Bucket(ctx context.Context, cfg config.Config) (*S3Bucket, error) {
	awsCfg, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithRegion(cfg.AWS.Region), // ensure region from your config
	)
	if err != nil {
		return nil, errors.Wrap(err, "[S3Bucket-NewS3Bucket] failed to load aws config")
	}

	return &S3Bucket{
		cfg:    cfg,
		client: s3.NewFromConfig(awsCfg),
	}, nil
}

// Upload uploads a file to the configured S3 bucket and returns its stored path.
func (s *S3Bucket) Upload(ctx context.Context, f *multipart.FileHeader, folder string) (string, error) {
	src, err := f.Open()
	if err != nil {
		return "", errors.Wrap(err, "[S3Bucket-Upload] unable to open file")
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	// Extract file extension
	ext := filepath.Ext(f.Filename)
	fileName := fmt.Sprintf("%s%s", helper.SHAEncrypt(f.Filename), ext)
	objectKey := fmt.Sprintf("%s/%s", folder, fileName)

	uploader := manager.NewUploader(s.client)

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.cfg.AWS.BucketName),
		Key:    aws.String(objectKey),
		Body:   src,
		ACL:    types.ObjectCannedACLPublicRead, // ✅ enum instead of *string
	})
	if err != nil {
		return "", fmt.Errorf("[S3Bucket-Upload] failed to upload file: %w", err)
	}

	return objectKey, nil
}

// Delete removes a file from the S3 bucket and waits until it's gone.
func (s *S3Bucket) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.cfg.AWS.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("[S3Bucket-Delete] failed to delete object: %w", err)
	}

	waiter := s3.NewObjectNotExistsWaiter(s.client)

	// This sets the *overall maximum wait time* (not just a timeout ctx).
	maxWait := 30 * time.Second

	if err := waiter.Wait(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.cfg.AWS.BucketName),
		Key:    aws.String(key),
	}, maxWait, func(o *s3.ObjectNotExistsWaiterOptions) {
		o.MinDelay = 2 * time.Second
		o.MaxDelay = 5 * time.Second
	}); err != nil {
		return fmt.Errorf("[S3Bucket-Delete] waiter failed: %w", err)
	}

	return nil
}
