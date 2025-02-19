package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type authClient struct {
	ClientId          string
	ClientSecret      string
	UserIdOrLoginName string
}

type accessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

const (
	Endpoint       string = "https://secure-wms.com/AuthServer/api/Token"
	Host           string = "secure-wms.com"
	Connection     string = "keep-alive"
	ContentType    string = "application/json"
	Accept         string = "application/json"
	AcceptEncoding string = "gzip,deflate,sdch"
	AcceptLanguage string = "en-US,en;q=0.8"
	GrantType      string = "client_credentials"
)

func New(clientId string, clientSecret string, userIdOrLoginName string) *authClient {

	return &authClient{
		ClientId:          clientId,
		ClientSecret:      clientSecret,
		UserIdOrLoginName: userIdOrLoginName,
	}

}

func (a *authClient) GetAccessToken(ctx context.Context) (*accessToken, error) {

	values := map[string]string{
		"grant_type":    GrantType,
		"user_login_id": a.UserIdOrLoginName,
	}

	spew.Dump(a.encodeAuth())

	payload, _ := json.Marshal(values)

	req, err := http.NewRequestWithContext(ctx, "POST", Endpoint, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Host", Host)
	req.Header.Set("Connection", Connection)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("Accept", Accept)
	req.Header.Set("Accept-Encoding", AcceptEncoding)
	req.Header.Set("Accept-Language", AcceptLanguage)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", a.encodeAuth()))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"http request to 'secure-wms.com' failed with a status of %d: %s",
			res.StatusCode,
			string(resBody),
		)
	}

	var accessToken accessToken
	err = json.Unmarshal(resBody, &accessToken)

	if err != nil {
		return nil, err
	}

	return &accessToken, nil

}

func (a *authClient) encodeAuth() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", a.ClientId, a.ClientSecret)))
}
