package models

type UploadImageRequest struct {
	Message MessageAttachment `json:"message"`
}

type MessageAttachment struct {
	Attachment AttachmentUpload `json:"attachment"`
}

type AttachmentUpload struct {
	Type    string            `json:"type"`
	Payload AttachmentPayload `json:"payload"`
}

type AttachmentPayload struct {
	IsReusable bool   `json:"is_reusable"`
	URL        string `json:"url"`
}

type UploadImageResponse struct {
	AttachmentID string `json:"attachment_id"`
	Error        struct {
		Message      string `json:"message"`
		Type         string `json:"type"`
		Code         int    `json:"code"`
		ErrorSubcode int    `json:"error_subcode"`
		FbtraceId    string `json:"fbtrace_id"`
	} `json:"error"`
}
