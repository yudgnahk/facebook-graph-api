package page

import (
	"fmt"
	httputils "github.com/yudgnahk/go-common-utils/http"
	"github.com/yudgnahk/gofacebook/constants"
	"github.com/yudgnahk/gofacebook/models"
	"github.com/yudgnahk/gofacebook/utils"
	"net/http"
)

func (c *pageClient) GetMessageByID(messageID string) (*models.GetMessageResponse, error) {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetMessageEndpoint, messageID), http.MethodGet)

	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	utils.AddJsonHeader(request)

	var response models.GetMessageResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
