package contents

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucket = "runestones"

func NewS3(ctx context.Context) *S3 {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	return &S3{client: s3.NewFromConfig(cfg)}
}

type S3 struct {
	client *s3.Client
}

func (st *S3) SaveContent(ctx context.Context, key string, content []byte) error {
	expirationTime := time.Now().Add(24 * time.Hour * 7)
	_, err := st.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:  aws.String(bucket),
		Key:     aws.String(generateKey(key)),
		Body:    bytes.NewReader(content),
		Expires: &expirationTime,
	})
	return err
}

func (st *S3) FindContent(ctx context.Context, key string) ([]byte, error) {
	out, err := st.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(generateKey(key)),
	})
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(out.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func generateKey(key string) string {
	return key + ".txt"
}
