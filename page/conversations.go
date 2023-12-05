package page

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yudgnahk/gofacebook/constants"
	"github.com/yudgnahk/gofacebook/models"

	httputils "github.com/yudgnahk/go-common-utils/http"
	"github.com/yudgnahk/gofacebook/utils"
)

func (c *pageClient) GetConversations(limit *int, offset *int, before, after *string,
	fields ...models.GetConversationsFields) (*models.GetConversationsResponse, error) {

	url := ""
	if len(fields) > 0 {
		url = "me/conversations?"
		paramStr := "fields=" + utils.SliceToString(fields)
		url += paramStr
	} else {
		url = constants.GetConversationsEndpoint
	}

	if limit != nil {
		url += "&limit=" + fmt.Sprintf("%v", *limit)
	}

	if offset != nil {
		url += "&offset=" + fmt.Sprintf("%v", *offset)
	}

	if before != nil && after != nil {
		return nil, errors.New("before and after can't be used together")
	}

	if before != nil {
		url += "&before=" + *before
	} else if after != nil {
		url += "&after=" + *after
	}

	url = c.PrepareUrl(url, http.MethodGet)

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
	url := c.PrepareUrl(constants.GetConversationsEndpoint, http.MethodGet)
	return url
}

func (c *pageClient) GetConversationByID(conversationID string, before, after *string) (*models.GetConversationResponse, error) {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetConversationEndpoint, conversationID), http.MethodGet)

	if before != nil && after != nil {
		return nil, errors.New("before and after can't be used together")
	}

	if before != nil {
		url += "&before=" + *before
	}

	if after != nil {
		url += "&after=" + *after
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

func (c *pageClient) GetBasicConversationDataByID(conversationID string) (*models.GetBasicConversationDataResponse, error) {
	url := c.PrepareUrl(fmt.Sprintf(constants.GetBasicConversationDataEndpoint, conversationID), http.MethodGet)

	request, err := httputils.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var response models.GetBasicConversationDataResponse
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
