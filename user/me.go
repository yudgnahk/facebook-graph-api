package user

import (
	"github.com/yudgnahk/gofacebook/constants"
	"github.com/yudgnahk/gofacebook/models"
	"net/http"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

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
