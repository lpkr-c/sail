package cloud

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devinmcgloin/sail/pkg/slog"
)

const (
	bucketName = "sail-content"
)

// Upload stores the given sketch in DO Spaces
func Upload(sketch *bytes.Buffer, path string) error {

	config := &aws.Config{Region: aws.String("us-east-1")}

	sess, err := session.NewSession(config)
	if err != nil {
		slog.ErrorPrintf("error while constructing new aws session %s", err)
		return err
	}
	svc := s3.New(sess)

	fileBytes := bytes.NewReader(sketch.Bytes())

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(int64(sketch.Len())),
		ContentType:   aws.String("image/png"),
	}

	if err != nil {
		slog.ErrorPrintf("Error while creating AWS params %s", err)
		return err
	}

	_, err = svc.PutObject(params)
	if err != nil {
		slog.ErrorPrintf("Error while uploading to aws %s", err)
		return err
	}

	return nil
}
