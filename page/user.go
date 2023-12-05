package page

import (
	"fmt"
	"github.com/yudgnahk/gofacebook/constants"
	"github.com/yudgnahk/gofacebook/models"
	"net/http"
	"strings"

	httputils "github.com/yudgnahk/go-common-utils/http"
	"github.com/yudgnahk/gofacebook/utils"
)

func (c *pageClient) GetUser(userID string, fields ...models.GetUserField) (*models.GetUserResponse, error) {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetUserEndpoint, userID), http.MethodGet)

	if len(fields) > 0 {
		paramStr := "fields=" + utils.SliceToString(fields)

		if strings.Contains(url, "?") {
			url += "&"
		} else {
			url += "?"
		}

		url += paramStr
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetUserResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
