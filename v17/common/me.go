package common

import (
	"github.com/yudgnahk/facebook-graph-api/v17/constants"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
	httputils "github.com/yudgnahk/go-common-utils/http"
	"net/http"
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
