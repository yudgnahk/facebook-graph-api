package page

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yudgnahk/facebook-graph-api/utils"
	"github.com/yudgnahk/facebook-graph-api/v17/constants"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
	httputils "github.com/yudgnahk/go-common-utils/http"
)

func (c *pageClient) GetConversations(fields ...models.GetConversationsFields) (*models.GetConversationsResponse, error) {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetConversationsEndpoint, c.UserID), http.MethodGet)

	if len(fields) > 0 {
		paramStr := "fields=" + utils.SliceToString(fields)

		if strings.Contains(url, "?") {
			url += "&"
		} else {
			url += "?"
		}

		url += paramStr
	}

	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(request.URL.String())

	var response models.GetConversationsResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *pageClient) GetConversation(conversationID string, fields ...models.GetConversationFields) (*models.GetConversationResponse, error) {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetConversationEndpoint, conversationID), http.MethodGet)

	if len(fields) > 0 {
		paramStr := "fields=" + utils.SliceToString(fields)

		if strings.Contains(url, "?") {
			url += "&"
		} else {
			url += "?"
		}

		url += paramStr
	}

	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(request.URL.String())

	var response models.GetConversationResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
