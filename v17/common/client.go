package common

import (
	"net/http"
	"strings"

	"github.com/yudgnahk/facebook-graph-api/v17/constants"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
)

// Client is a client for the personal facebook API.
type Client struct {
	AccessToken              string
	UserID                   string
	ClientID                 string
	ClientSecret             string
	ExpiresIn                int
	DataAccessExpirationTime int
}

type IClient interface {
	GetMe() (*models.GetMeResponse, error)
	GetLongLivedToken() (*models.GetLongLivedTokenResponse, error)
	PrepareUrl(url, method string) string
}

// NewClient returns a new facebook API client.
func NewClient(accessToken, userID, clientID, clientSecret string, expiresIn, dataAccessExpirationTime int) Client {
	return Client{
		AccessToken:              accessToken,
		UserID:                   userID,
		ExpiresIn:                expiresIn,
		DataAccessExpirationTime: dataAccessExpirationTime,
		ClientID:                 clientID,
		ClientSecret:             clientSecret,
	}
}

func (c *Client) PrepareUrl(url, method string) string {
	finalURL := constants.BaseURL + url
	switch method {
	case http.MethodGet:
		// Check if url contains query string
		if strings.Contains(url, "?") {
			finalURL += "&"
		} else {
			finalURL += "?"
		}

		finalURL += "access_token=" + c.AccessToken + constants.ParametersForGetRequest
	}

	return finalURL
}
