package constants

const (
	// BaseURL is the base URL for the personal facebook API.
	BaseURL = "https://graph.facebook.com/v17.0/"

	GetMeEndpoint             = "me"
	GetAccountsEndpoint       = "me/accounts"
	GetLongLivedTokenEndpoint = "oauth/access_token"
	// GetConversationsEndpoint is the endpoint for getting conversations.
	GetConversationsEndpoint = "%v/conversations"
	GetConversationEndpoint  = "%v/messages"
	ParametersForGetRequest  = "&debug=all&format=json&method=get&pretty=0&suppress_http_code=1&transport=cors"
)
