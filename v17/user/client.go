package user

import (
	"github.com/yudgnahk/gofacebook/v17/common"
	"github.com/yudgnahk/gofacebook/v17/models"
)

// userClient is a client for the personal facebook API.
type userClient struct {
	common.Client
}

type UserClient interface {
	GetMe() (*models.GetMeResponse, error)
	GetAccounts() (*models.GetAccountsResponse, error)
	BatchRequest(requests []models.BatchRequest) ([]models.BatchResponse, error)
}

// NewUserClient returns a new facebook API client.
func NewUserClient(client common.Client) UserClient {
	return &userClient{
		client,
	}
}
