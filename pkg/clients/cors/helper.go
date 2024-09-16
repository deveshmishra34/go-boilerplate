package cors

import "github.com/deveshmishra34/groot/pkg/utils/constants"

var client *CorsClient

func init() {
	client = &CorsClient{
		name: constants.FEATURE_CORS,
	}
}

func GetClient() *CorsClient {
	return client
}
