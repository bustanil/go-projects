package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	bucketName    string
	presignClient *s3.PresignClient
}

func (c *Client) PresignPutObject(ctx context.Context, path string) (string, map[string][]string, error) {
	req, err := c.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &c.bucketName,
		Key:    &path,
	})
	if err != nil {
		return "", nil, err
	}

	return req.URL, req.SignedHeader, nil
}

func NewClient(cfg *aws.Config, bucketName string) *Client {
	c := s3.NewFromConfig(*cfg)
	presignClient := s3.NewPresignClient(c)
	return &Client{
		bucketName:    bucketName,
		presignClient: presignClient,
	}
}
