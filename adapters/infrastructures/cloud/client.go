package cloud

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	Client *s3.Client
}

func NewS3Client() (*S3, error) {
	cfgs, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &S3{
		Client: s3.NewFromConfig(cfgs, func(o *s3.Options) {
			o.HTTPClient = &http.Client{
				Transport: &http.Transport{
					ResponseHeaderTimeout: 15 * time.Minute,
					Proxy:                 http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						KeepAlive: 10,
						Timeout:   2 * time.Second,
					}).DialContext,
					MaxIdleConns:        10,
					IdleConnTimeout:     30 * time.Second,
					TLSHandshakeTimeout: 5 * time.Second,
					MaxIdleConnsPerHost: 10,
				},
			}
		}),
	}, nil
}
