package casdoor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type client struct {
	casdoor *casdoorsdk.Client
}

func newClient(id, secret string) *client {
	return &client{
		casdoor: casdoorsdk.NewClient(
			"http://auth:8000",
			id,
			secret,
			"certificate",
			"built-in",
			"app-built-in",
		),
	}
}

func (c *client) getAccess(code, state string) (string, error) {
	token, err := c.casdoor.GetOAuthToken(code, state)
	if err != nil {
		return "", err
	}

	if token == nil {
		return "", errors.New("empty token")
	}

	return token.AccessToken, nil
}

func (c *client) getUserInfo(token string) (AuthUser, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://auth:8000/api/userinfo",
		http.NoBody,
	)
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

	if response.StatusCode != http.StatusOK {
		return AuthUser{}, fmt.Errorf("casdoor userinfo status: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return AuthUser{}, err
	}

	var user AuthUser

	err = json.Unmarshal(body, &user)
	if err != nil {
		return AuthUser{}, err
	}

	return user, nil
}
