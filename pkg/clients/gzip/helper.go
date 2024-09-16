package gzip

import "github.com/deveshmishra34/groot/pkg/utils/constants"

var client *GzipClient

func init() {
	client = &GzipClient{
		name: constants.FEATURE_GZIP,
	}
}

func GetClient() *GzipClient {
	return client
}
