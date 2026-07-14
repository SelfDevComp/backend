package casdoor

import (
	"encoding/json"
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type client struct {
	casdoor *casdoorsdk.Client
}

func newClient(id, secret string) *client {
	return &client{
		casdoor: casdoorsdk.NewClient("endpoint", id, secret, "certificate", "my-org", "my-app"),
	}
}

func (c *client) getAccess(code, state string) (string, error) {
	token, err := casdoorsdk.GetOAuthToken(code, state)
	return token.AccessToken, err
}

func (c *client) getUserInfo(token string) (AuthUser, error) {
	request, err := http.NewRequest(http.MethodGet, "http://auth.self-dev.test/api/userinfo", http.NoBody)
	if err != nil {
		return AuthUser{}, err
	}

	request.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return AuthUser{}, err
	}
	defer response.Body.Close()

	var user AuthUser

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return AuthUser{}, err
	}

	return user, nil
}
