package page

import (
	"github.com/yudgnahk/facebook-graph-api/v17/common"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
)

type PageClient interface {
	GetConversations(fields ...models.GetConversationsFields) (*models.GetConversationsResponse, error)
	GetConversationByID(conversationID string, fields ...models.GetConversationFields) (*models.GetConversationResponse, error)
	GetUser(userID string, fields ...models.GetUserField) (*models.GetUserResponse, error)
	BatchRequest(requests []models.BatchRequest) ([]models.BatchResponse, error)
	GetConversationsEndpoint() string
	SendMessage(recipientID string, message string) error
}

type pageClient struct {
	common.Client
}

func NewPageClient(client common.Client) PageClient {
	return &pageClient{
		client,
	}
}
