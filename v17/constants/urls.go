package constants

const (
	// BaseURL is the base URL for the personal facebook API.
	BaseURL = "https://graph.facebook.com/v17.0/"

	GetMeEndpoint             = "me"
	GetAccountsEndpoint       = "me/accounts"
	GetLongLivedTokenEndpoint = "oauth/access_token"

	GetUserEndpoint = "%v"

	// GetConversationsEndpoint is the endpoint for getting conversations.
	GetConversationsEndpoint = "%v/conversations?fields=message_count,participants,unread_count,messages.limit(1){created_time,message,sticker,attachments{image_data,video_data,file_url}}"
	GetConversationEndpoint  = "%v/messages?fields=message,created_time,id,from,sticker,attachments{image_data,file_url,video_data,name,size}&limit=20"
	ParametersForGetRequest  = "&debug=all&format=json&method=get&pretty=0&suppress_http_code=1&transport=cors"
)
