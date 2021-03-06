package ciossdk

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func (self Auth) GetAccessTokenByRefreshToken() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
	if self.debug {
		log.Printf("%s", "Refresh AccessToken.")
	}
	response, _, err := self.ApiClient.AuthApi.
		RefreshToken(context.Background()).
		GrantType("refresh_token").
		RefreshToken(self.ref).
		ClientId(self.clientId).
		ClientSecret(self.clientSecret).
		Scope(self.scope).
		Execute()
	if err != nil {
		return "", "", "", 0, err
	}
	return response.AccessToken, response.Scope, response.TokenType, int(response.ExpiresIn), nil

}

func (self Auth) GetAccessTokenOnClient() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
	responseData := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}
	values := url.Values{}
	values.Add("grant_type", "client_credentials")
	values.Add("client_id", self.clientId)
	values.Add("client_secret", self.clientSecret)
	values.Add("scope", self.scope)
	resp, _ := http.Post(
		self.Url+"/connect/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(values.Encode()),
	)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", 0, err
	}

	err = convert.UnMarshalJson(body, &responseData)
	if err != nil {
		return "", "", "", 0, err
	}
	return responseData.AccessToken, responseData.Scope, responseData.TokenType, responseData.ExpiresIn, nil
}

func (self Auth) GetAccessTokenOnDevice() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
	return "", "", "", 0, errors.New("not implement")

}
