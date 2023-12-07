package constants

const (
	// BaseURL is the base URL for the personal facebook API.
	BaseURL = "https://graph.facebook.com/v18.0/"

	GetMeEndpoint             = "me"
	GetAccountsEndpoint       = "me/accounts"
	GetLongLivedTokenEndpoint = "oauth/access_token"

	GetUserEndpoint = "%v"

	// GetConversationsEndpoint is the endpoint for getting conversations.
	GetConversationsEndpoint = "me/conversations?fields=message_count,participants,updated_time,unread_count,messages.limit(1){created_time,message,sticker,attachments{image_data,video_data,file_url},from,id}"
	GetConversationEndpoint  = "%v/messages?fields=message,created_time,id,from,sticker,attachments{image_data,file_url,video_data,name,size}&limit=20"
	GetMessageEndpoint       = "%v?fields=message,created_time,id,from,sticker,attachments{image_data,file_url,video_data,name,size}"

	GetBasicConversationDataEndpoint = "%v?fields=participants,id,message_count,unread_count,updated_time"
	SendMessageEndpoint              = "me/messages"
	ParametersForGetRequest          = "&debug=all&format=json&method=get&pretty=0&suppress_http_code=1&transport=cors"
)
