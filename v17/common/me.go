package common

import (
	"net/http"

	httputils "github.com/yudgnahk/go-common-utils/http"
	"github.com/yudgnahk/gofacebook/v17/constants"
	"github.com/yudgnahk/gofacebook/v17/models"
)

func (c *Client) GetMe() (*models.GetMeResponse, error) {
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
