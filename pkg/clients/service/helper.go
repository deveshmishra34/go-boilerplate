package service

import "github.com/deveshmishra34/groot/pkg/utils/constants"

var client *ServiceClient

func init() {
	client = &ServiceClient{
		name: constants.FEATURE_SERVICE,
	}
}

func GetClient() *ServiceClient {
	return client
}
