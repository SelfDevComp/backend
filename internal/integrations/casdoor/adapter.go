package casdoor

import (
	"github.com/sklyar-vlad/selfDev/internal/config"
)

type adapter struct {
	client *client
}

func NewAdapter(cfg config.ConfigAuth) *adapter {
	return &adapter{
		client: newClient(cfg.ClientId, cfg.ClientSecret),
	}
}

func (a *adapter) GetAccess(code, state string) (string, error) {
	return a.client.getAccess(code, state)
}

func (a *adapter) GetUserInfo(token string) (AuthUser, error) {
	return a.client.getUserInfo(token)
}
