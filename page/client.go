package page

import (
	"github.com/yudgnahk/gofacebook/common"
	"github.com/yudgnahk/gofacebook/models"
)

type PageClient interface {
	GetConversations(limit *int, offset *int, before, after *string,
		fields ...models.GetConversationsFields) (*models.GetConversationsResponse, error)
	GetConversationByID(conversationID string, before, after *string) (*models.GetConversationResponse, error)
	GetUser(userID string, fields ...models.GetUserField) (*models.GetUserResponse, error)
	BatchRequest(requests []models.BatchRequest) ([]models.BatchResponse, error)
	GetConversationsEndpoint() string
	SendMessage(recipientID string, message string) (*models.GetMessageResponse, error)
	SendAttachment(recipientID string, attachmentID string) (*models.GetMessageResponse, error)
	GetBasicConversationDataByID(conversationID string) (*models.GetBasicConversationDataResponse, error)
	UploadImageByUrl(url string) (*models.UploadImageResponse, error)
	GetMessageByID(messageID string) (*models.GetMessageResponse, error)
}

type pageClient struct {
	common.Client
}

func NewPageClient(client common.Client) PageClient {
	return &pageClient{
		client,
	}
}
