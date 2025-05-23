package data

import (
	"agents/app/authn/service/internal/biz"
	"agents/app/authn/service/internal/conf"
	"context"
	"encoding/json"
	"fmt"
)

var _ biz.GatewayRepo = (*gateway)(nil)

type gateway struct {
	data *Data
	kong *conf.Kong
}

func NewGatewayRepo(data *Data, kong *conf.Kong) biz.GatewayRepo {
	return &gateway{
		data: data,
		kong: kong,
	}
}

// CreateUserConsumer implements biz.Gateway.
func (g *gateway) CreateUserConsumer(ctx context.Context, user *biz.User) (string, error) {
	role := "user"
	consumerUsername := fmt.Sprintf("%s-%s", role, user.Username)

	req := g.data.cli.R().SetContext(ctx)
	req.SetBody(map[string]any{
		"custom_id": user.Id,
		"username":  consumerUsername,
		"tags": []string{
			role,
		},
	})

	if g.kong.ApiKey != "" {
		req.SetHeader("X-API-KEY", g.kong.ApiKey)
	}

	resp, err := req.Post(g.kong.AdminApi + "/consumers")
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("%s", string(resp.Body()))
	}

	return consumerUsername, nil
}

// EnableJwtPluginForConsumer implements biz.Gateway.
func (g *gateway) CreateConsumerCredential(ctx context.Context, consumer string, pubkeyPem string) (string, error) {
	req := g.data.cli.R().SetContext(ctx)
	req.SetBody(map[string]string{
		"algorithm":      biz.DefaultSigningAlg,
		"rsa_public_key": pubkeyPem,
	})
	req.SetPathParam("consumer", consumer)

	if g.kong.ApiKey != "" {
		req.SetHeader("X-API-KEY", g.kong.ApiKey)
	}

	resp, err := req.Post(g.kong.AdminApi + "/consumers/{consumer}/jwt")
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("%s", string(resp.Body()))
	}

	var reply struct {
		Key string `json:"key"`
	}

	err = json.Unmarshal(resp.Body(), &reply)
	if err != nil {
		return "", err
	}

	return reply.Key, nil
}
