package common

import (
	"net/http"

	"github.com/yudgnahk/facebook-graph-api/v17/constants"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
	httputils "github.com/yudgnahk/go-common-utils/http"
)

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
