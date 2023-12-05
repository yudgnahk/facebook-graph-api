package common

import (
	"bytes"
	"encoding/json"
	"github.com/yudgnahk/gofacebook/models"
	"net/http"

	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *Client) BatchRequest(requests []models.BatchRequest) ([]models.BatchResponse, error) {
	url := c.PrepareUrl("", http.MethodPost)

	body := map[string]interface{}{
		"batch": requests,
	}

	batchRequestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := httputils.NewRequest(http.MethodPost, url, bytes.NewBuffer(batchRequestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+c.AccessToken)

	var response []models.BatchResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
