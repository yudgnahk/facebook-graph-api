package models

type GetMeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetAccountsResponse struct {
	Data []struct {
		Picture struct {
			Data struct {
				Height       int    `json:"height"`
				IsSilhouette bool   `json:"is_silhouette"`
				Url          string `json:"url"`
				Width        int    `json:"width"`
			} `json:"data"`
		} `json:"picture,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		AccessToken string `json:"access_token,omitempty"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}

type GetLongLivedTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Error       struct {
		Message      string `json:"message"`
		Type         string `json:"type"`
		Code         int    `json:"code"`
		ErrorSubcode int    `json:"error_subcode"`
		FbtraceId    string `json:"fbtrace_id"`
	} `json:"error"`
}

type BatchRequest struct {
	Method      string `json:"method"`
	RelativeUrl string `json:"relative_url"`
}

type BatchResponse struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}
