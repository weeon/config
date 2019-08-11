package config

import (
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/weeon/contract"
)

type ConsulConfig struct {
	client *ConsulClient
}

var _ contract.Config = new(ConsulConfig)

func (c *ConsulConfig) Get(key string) ([]byte, error) {
	return c.client.KVGet(key)
}

func (c *ConsulConfig) Set(key string, value []byte) error {
	return c.client.KVSet(key, value)
}

func NewConsulConfig(host, token string) (*ConsulConfig, error) {

	client, err := NewConsulClient(&api.Config{
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

func NewConsulConfigFromEnv() (*ConsulConfig, error) {
	consulAddr := os.Getenv("CONSUL_ADDR")
	consulToken := os.Getenv("CONSUL_TOKEN")
	return NewConsulConfig(consulAddr, consulToken)
}
