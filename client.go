package config

import "github.com/hashicorp/consul/api"

type ConsulClient struct {
	Client *api.Client
}

func NewConsulClient(cfg *api.Config) (*ConsulClient, error) {
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsulClient{
		Client: client,
	}, nil
}

func (c *ConsulClient) KVGet(key string) ([]byte, error) {
	p, _, err := c.Client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return []byte{}, nil
	}
	return p.Value, nil
}

func (c *ConsulClient) KVSet(key string, value []byte) error {
	p := &api.KVPair{Key: key, Value: value}
	_, err := c.Client.KV().Put(p, nil)
	return err
}
