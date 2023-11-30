package page

import (
	"bytes"
	"encoding/json"
	"errors"
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

	var response models.GetConversationsResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *pageClient) GetConversationsEndpoint() string {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetConversationsEndpoint, c.UserID), http.MethodGet)
	return url
}

func (c *pageClient) GetConversationByID(conversationID string, fields ...models.GetConversationFields) (*models.GetConversationResponse, error) {
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

	var response models.GetConversationResponse
	err = httputils.Execute(request, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *pageClient) SendMessage(recipientID string, message string) error {
	url := c.PrepareUrl(constants.SendMessageEndpoint, http.MethodPost)

	var req models.SendMessageRequest
	req.Recipient.ID = recipientID
	req.Message.Text = message

	if len(message) == 0 {
		return errors.New("message can't be empty")
	}

	// marshal request data
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshall request: %w", err)
	}

	request, err := httputils.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")

	var response interface{}
	err = httputils.Execute(request, &response)
	if err != nil {
		return err
	}

	return nil
}
