package config

import (
	"github.com/weeon/contract"
	"github.com/hashicorp/consul/api"
	"github.com/orvice/kit/consul"
)

type ConsulConfig struct {
	client *consul.Client
}

var _ contract.Config = new(ConsulConfig)

func (c *ConsulConfig) Get(key string) ([]byte, error) {
	return c.client.KVGet(key)
}

func (c *ConsulConfig) Set(key string, value []byte) error {
	return c.client.KVSet(key, value)
}

func NewConsulConfig(host, token string) (*ConsulConfig, error) {

	client, err := consul.NewClient(&api.Config{
		Address: host,
		Token:   token,
	})
	if err != nil {
		return nil, err
	}

	return &ConsulConfig{
		client: client,
	}, nil
}
