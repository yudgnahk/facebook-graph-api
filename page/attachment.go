package page

import (
	"bytes"
	"encoding/json"
	httputils "github.com/yudgnahk/go-common-utils/http"
	"github.com/yudgnahk/gofacebook/models"
	"github.com/yudgnahk/gofacebook/utils"
	"net/http"
)

func (c *pageClient) UploadImageByUrl(imageUrl string) (*models.UploadImageResponse, error) {
	url := c.PrepareUrl("me/message_attachments", http.MethodPost)

	req := models.UploadImageRequest{
		Message: models.MessageAttachment{
			Attachment: models.AttachmentUpload{
				Type: "image",
				Payload: models.AttachmentPayload{
					IsReusable: true,
					URL:        imageUrl,
				},
			},
		},
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := httputils.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	utils.AddJsonHeader(request)

	var response models.UploadImageResponse
	err = httputils.Execute(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
