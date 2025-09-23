// Package gcs provides function for gcs products
package gcs

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"

	"gin-starter/common/helper"
	"gin-starter/config"
)

// GoogleCloudStorage define struct for gcs integration
type GoogleCloudStorage struct {
	cfg config.Config
}

// NewGoogleCloudStorage initiate gcs sdk
func NewGoogleCloudStorage(cfg config.Config) *GoogleCloudStorage {
	return &GoogleCloudStorage{cfg: cfg}
}

// Upload uploads file to bucket
func (g *GoogleCloudStorage) Upload(f *multipart.FileHeader, folder string) (string, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_SA")))

	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error get config json")
	}

	src, err := f.Open()
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error open file")
	}

	defer func() {
		if err := src.Close(); err != nil {
			fmt.Println("error while closing gin context form file :", err)
		}
	}()

	splitFilename := strings.Split(f.Filename, ".")
	fileMime := ""

	if len(splitFilename) > 0 {
		fileMime = splitFilename[len(splitFilename)-1]
	}

	fileName := fmt.Sprintf("%s.%s", helper.SHAEncrypt(f.Filename), fileMime)

	fileStored := fmt.Sprintf("%s/%s", folder, fileName)

	sw := storageClient.Bucket(bucket).Object(fileStored).NewWriter(ctx)

	if _, err = io.Copy(sw, src); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error copy file")
	}

	if err := sw.Close(); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error close file")
	}

	u, err := url.Parse("/" + sw.Attrs().Name)
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error parse url")
	}

	// Make public
	acl := storageClient.Bucket(bucket).Object(fileStored).ACL()

	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-HandleFileUploadToBucket] error while making object public")
	}

	return u.String(), nil
}

// UploadSavedFile uploads file to bucket
func (g *GoogleCloudStorage) UploadSavedFile(filepath, folder string) (string, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_SA")))

	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error get config json")
	}

	file, err := os.Open(filepath) // #nosec
	if err != nil {
		return "", errors.Wrap(err, "problem opening file for gcs")
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error while closing context file :", err)
		}
	}()

	splitFilename := strings.Split(file.Name(), ".")
	fileMime := ""

	if len(splitFilename) > 0 {
		fileMime = splitFilename[len(splitFilename)-1]
	}

	fileName := fmt.Sprintf("%s.%s", helper.SHAEncrypt(file.Name()), fileMime)

	fileStored := fmt.Sprintf("%s/%s", folder, fileName)

	sw := storageClient.Bucket(bucket).Object(fileStored).NewWriter(ctx)

	if _, err = io.Copy(sw, file); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error copy file")
	}

	if err := sw.Close(); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error close file")
	}

	u, err := url.Parse("/" + sw.Attrs().Name)
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error parse url")
	}

	// Make public
	acl := storageClient.Bucket(bucket).Object(fileStored).ACL()

	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-HandleFileUploadToBucket] error while making object public")
	}

	return u.String(), nil
}

// Delete delete file from bucket
func (g *GoogleCloudStorage) Delete(path string) error {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_SA")))

	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Delete] error get config json")
	}

	if err := storageClient.Bucket(bucket).Object(url.QueryEscape(path)).Delete(context.Background()); err != nil {
		return errors.Wrap(err, fmt.Sprintf("[CloudStorageService-Delete] unable to delete bucket %q, file %q", bucket, url.QueryEscape(path)))
	}

	return nil
}
