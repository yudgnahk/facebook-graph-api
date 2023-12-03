package models

type GetConversationsResponse struct {
	Data []struct {
		MessageCount int `json:"message_count"`
		Participants struct {
			Data []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"participants"`
		UnreadCount int `json:"unread_count"`
		Messages    struct {
			Data []struct {
				CreatedTime string `json:"created_time"`
				Message     string `json:"message,omitempty"`
				From        struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"from,omitempty"`
				Attachments struct {
					Data []struct {
						ImageData struct {
							Width           int    `json:"width,omitempty"`
							Height          int    `json:"height,omitempty"`
							MaxWidth        int    `json:"max_width,omitempty"`
							MaxHeight       int    `json:"max_height,omitempty"`
							URL             string `json:"url,omitempty"`
							PreviewURL      string `json:"preview_url,omitempty"`
							ImageType       int    `json:"image_type,omitempty"`
							RenderAsSticker bool   `json:"render_as_sticker,omitempty"`
						} `json:"image_data,omitempty"`
						VideoData struct {
							Width      int    `json:"width,omitempty"`
							Height     int    `json:"height,omitempty"`
							Length     int    `json:"length,omitempty"`
							VideoType  int    `json:"video_type,omitempty"`
							URL        string `json:"url,omitempty"`
							PreviewURL string `json:"preview_url,omitempty"`
							Rotation   int    `json:"rotation,omitempty"`
						} `json:"video_data,omitempty"`
						FileURL string `json:"file_url,omitempty"`
						ID      string `json:"id,omitempty"`
					} `json:"data,omitempty"`
					Paging struct {
						Cursors struct {
							Before string `json:"before,omitempty"`
							After  string `json:"after,omitempty"`
						} `json:"cursors,omitempty"`
						Next string `json:"next,omitempty"`
					} `json:"paging,omitempty"`
				} `json:"attachments,omitempty"`
				ID      string `json:"id"`
				Sticker string `json:"sticker,omitempty"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					Before string `json:"before,omitempty"`
					After  string `json:"after,omitempty"`
				} `json:"cursors,omitempty"`
				Next string `json:"next,omitempty"`
			} `json:"paging,omitempty"`
		} `json:"messages"`
		ID string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before,omitempty"`
			After  string `json:"after,omitempty"`
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type GetConversationResponse struct {
	Data []struct {
		From struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Id    string `json:"id"`
		} `json:"from"`
		Message     string `json:"message"`
		Attachments struct {
			Data []struct {
				FileUrl   string `json:"file_url,omitempty"`
				Id        string `json:"id"`
				ImageData struct {
					Width           int    `json:"width"`
					Height          int    `json:"height"`
					MaxWidth        int    `json:"max_width"`
					MaxHeight       int    `json:"max_height"`
					Url             string `json:"url"`
					PreviewUrl      string `json:"preview_url"`
					ImageType       int    `json:"image_type"`
					RenderAsSticker bool   `json:"render_as_sticker"`
				} `json:"image_data,omitempty"`
				VideoData struct {
					Width      int    `json:"width,omitempty"`
					Height     int    `json:"height,omitempty"`
					Length     int    `json:"length,omitempty"`
					VideoType  int    `json:"video_type,omitempty"`
					URL        string `json:"url,omitempty"`
					PreviewURL string `json:"preview_url,omitempty"`
					Rotation   int    `json:"rotation,omitempty"`
				} `json:"video_data,omitempty"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					Before string `json:"before"`
					After  string `json:"after"`
				} `json:"cursors"`
				Next string `json:"next"`
			} `json:"paging"`
		} `json:"attachments,omitempty"`
		Id      string `json:"id"`
		Sticker string `json:"sticker,omitempty"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

type GetBasicConversationDataResponse struct {
	Participants struct {
		Data []struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Id    string `json:"id"`
		} `json:"data"`
	} `json:"participants"`
	UpdatedTime  string `json:"updated_time"`
	Id           string `json:"id"`
	MessageCount int    `json:"message_count"`
	UnreadCount  int    `json:"unread_count"`
}

type GetConversationsFields string

const (
	CanReplyGetConversationsField     GetConversationsFields = "can_reply"
	IsSubscribedGetConversationsField GetConversationsFields = "is_subscribed"
	MessageCountGetConversationsField GetConversationsFields = "message_count"
	UnreadCountGetConversationsField  GetConversationsFields = "unread_count"
	SubjectGetConversationsField      GetConversationsFields = "subject"
	ParticipantsGetConversationsField GetConversationsFields = "participants"
	UpdatedTimeGetConversationsField  GetConversationsFields = "updated_time"
	MessagesGetConversationsField     GetConversationsFields = "messages.limit(1){message}"
)

type GetConversationFields string

const (
	MessageGetConversationField     GetConversationFields = "message"
	FromGetConversationField        GetConversationFields = "from"
	StickerGetConversationField     GetConversationFields = "sticker"
	CreatedTimeGetConversationField GetConversationFields = "created_time"
	IDGetConversationField          GetConversationFields = "id"
)

type SendMessageRequest struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}
