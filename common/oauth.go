package common

import (
	"net/http"

	"github.com/yudgnahk/gofacebook/constants"
	"github.com/yudgnahk/gofacebook/models"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *Client) GetLongLivedTokenURL() string {
	return constants.GetLongLivedTokenEndpoint +
		"?grant_type=fb_exchange_token&client_id=" +
		c.ClientID +
		"&client_secret=" +
		c.ClientSecret +
		"&fb_exchange_token=" +
		c.AccessToken
}

// use this for batch request to get long-lived token of many pages
func (c *Client) GetLongLivedTokenURLWithToken(token string) string {
	return constants.GetLongLivedTokenEndpoint +
		"?grant_type=fb_exchange_token&client_id=" +
		c.ClientID +
		"&client_secret=" +
		c.ClientSecret +
		"&fb_exchange_token=" +
		token
}

func (c *Client) GetLongLivedToken() (*models.GetLongLivedTokenResponse, error) {
	url := constants.BaseURL +
		constants.GetLongLivedTokenEndpoint +
		"?grant_type=fb_exchange_token&client_id=" +
		c.ClientID +
		"&client_secret=" +
		c.ClientSecret +
		"&fb_exchange_token=" +
		c.AccessToken

	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetLongLivedTokenResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
