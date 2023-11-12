package models

type GetConversationsResponse struct {
	Data []struct {
		MessageCount int `json:"message_count"`
		UnreadCount  int `json:"unread_count"`
		Participants struct {
			Data []struct {
				Name  string `json:"name"`
				Email string `json:"email"`
				Id    string `json:"id"`
			} `json:"data"`
		} `json:"participants"`
		UpdatedTime string `json:"updated_time"`
		Messages    struct {
			Data []struct {
				Message string `json:"message"`
				Id      string `json:"id"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					Before string `json:"before"`
					After  string `json:"after"`
				} `json:"cursors"`
				Next string `json:"next"`
			} `json:"paging"`
		} `json:"messages"`
		Id string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

type GetConversationResponse struct {
	Data []struct {
		Message string `json:"message"`
		From    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Id    string `json:"id"`
		} `json:"from"`
		To struct {
			Data []struct {
				Name  string `json:"name"`
				Email string `json:"email"`
				Id    string `json:"id"`
			} `json:"data"`
		} `json:"to"`
		Id string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
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
