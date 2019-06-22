package dumper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Dumper
type Dumper interface {
	Dump(ctx context.Context, r io.Reader, path string) error
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Uploader
type Uploader interface {
	Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

type RoundRobinDumpClient struct {
	logger  *zap.Logger
	clients []Dumper
}

func NewRoundRobinDumpClient(logger *zap.Logger, clients ...Dumper) RoundRobinDumpClient {
	return RoundRobinDumpClient{
		logger,
		clients,
	}
}

func (c RoundRobinDumpClient) Dump(ctx context.Context, r io.Reader, path string) error {
	// this could potentially load a lot into memory, but we have to buffer it somehow so that we can read it into multiple
	// dump clients.  This could potentially be better if we use Go pipelining here, but for now i'm keeping it as is
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	for _, client := range c.clients {
		br := bytes.NewReader(b)
		err := client.Dump(ctx, br, path)
		if err != nil {
			return err
		}
	}
	return err
}

type LocalDumpHandler struct {
	path   string
	logger *zap.Logger
	fs     afero.Fs
}

func NewLocalDumpHandler(path string, logger *zap.Logger, fs afero.Fs) LocalDumpHandler {
	return LocalDumpHandler{
		path,
		logger,
		fs,
	}
}

func (c LocalDumpHandler) Dump(ctx context.Context, r io.Reader, path string) error {
	c.logger.Debug(fmt.Sprintf("Local dump to %s", path))
	location := filepath.Join(c.path, path)

	f, err := c.fs.Create(location)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	return f.Close()
}

func NewS3DumpHandler(uploader Uploader, bucket string, logger *zap.Logger) S3DumpHandler {
	return S3DumpHandler{
		uploader,
		bucket,
		logger,
	}
}

type S3DumpHandler struct {
	uploader Uploader
	bucket   string
	logger   *zap.Logger
}

func (c S3DumpHandler) Dump(ctx context.Context, r io.Reader, path string) error {
	c.logger.Debug(fmt.Sprintf("S3 dump to bucket %s, path %s", c.bucket, path))
	_, err := c.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
		Body:   r,
	})
	return err
}
