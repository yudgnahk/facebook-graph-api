package page

import (
	"github.com/yudgnahk/facebook-graph-api/v17/common"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
)

type PageClient interface {
	GetConversations(fields ...models.GetConversationsFields) (*models.GetConversationsResponse, error)
	GetConversation(conversationID string, fields ...models.GetConversationFields) (*models.GetConversationResponse, error)
}

type pageClient struct {
	common.Client
}

func NewPageClient(client common.Client) PageClient {
	return &pageClient{
		client,
	}
}
