package user

import (
	"net/http"

	"github.com/yudgnahk/facebook-graph-api/v17/constants"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *userClient) GetMe() (*models.GetMeResponse, error) {
	url := c.PrepareUrl(constants.GetMeEndpoint, http.MethodGet)
	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetMeResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *userClient) GetAccounts() (*models.GetAccountsResponse, error) {
	url := c.PrepareUrl(constants.GetAccountsEndpoint, http.MethodGet)
	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetAccountsResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
