package models

type GetMeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetAccountsResponse struct {
	Data []struct {
		AccessToken  string `json:"access_token"`
		Category     string `json:"category"`
		CategoryList []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"category_list"`
		Name  string   `json:"name"`
		ID    string   `json:"id"`
		Tasks []string `json:"tasks"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

type GetLongLivedTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
