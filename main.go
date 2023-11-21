package main

import (
	"fmt"
	"os"

	"github.com/yudgnahk/facebook-graph-api/v17/common"
	"github.com/yudgnahk/facebook-graph-api/v17/models"
	"github.com/yudgnahk/facebook-graph-api/v17/page"
)

func main() {
	// Read .env
	//fbToken := os.Getenv("FB_TOKEN")
	//fbToken := os.Getenv("LONG_LIVED_TOKEN")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	pageToken := os.Getenv("PAGE_LONG_LIVED_TOKEN")
	pageID := os.Getenv("PAGE_ID")

	// Create client
	//client := user.NewClient(fbToken, "", clientID, clientSecret, 0, 0)
	//
	//// Get me
	//me, err := client.GetMe()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(err)
	//
	//// Print me
	//fmt.Println(me.ID, me.Name)
	//
	//// Get accounts
	//accountsResponse, err := client.GetAccounts()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, account := range accountsResponse.Data {
	//	fmt.Println(account.ID, account.Name, account.AccessToken)
	//}

	// Get long-lived token for the first page
	c1 := common.NewClient(pageToken, pageID, clientID, clientSecret, 0, 0)
	//longLivedTokenResponse, err := c1.GetLongLivedToken()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(longLivedTokenResponse.AccessToken)

	me, err := c1.GetMe()
	if err != nil {
		panic(err)
	}

	fmt.Println(me.ID, me.Name)

	pageClient := page.NewPageClient(c1)

	fields := []models.GetConversationsFields{
		models.MessageCountGetConversationsField,
		models.UnreadCountGetConversationsField,
		models.ParticipantsGetConversationsField,
		models.UpdatedTimeGetConversationsField,
		models.MessagesGetConversationsField,
	}

	conversations, err := pageClient.GetConversations(fields...)

	if err != nil {
		panic(err)
	}

	for _, conversation := range conversations.Data {
		// Get messages from each conversation

		conversationFields := []models.GetConversationFields{
			models.MessageGetConversationField,
			models.StickerGetConversationField,
			models.CreatedTimeGetConversationField,
			models.FromGetConversationField,
			models.IDGetConversationField,
		}

		messages, err := pageClient.GetConversationByID(conversation.Id, conversationFields...)
		if err != nil {
			panic(err)
		}

		fmt.Println(messages.Data[0].Message)
	}

	//Get Duy Khang user profile
	userFields := []models.GetUserField{
		models.IDGetUserField,
		models.NameGetUserField,
		models.PictureGetUserField.WithAttributes(map[models.PictureAttribute]string{
			models.WidthPictureAttribute:  "1000",
			models.HeightPictureAttribute: "1000",
		}),
	}

	kelvin, err := pageClient.GetUser("100000676495118", userFields...)
	if err != nil {
		panic(err)
	}

	fmt.Println(kelvin.Picture.Data.Url)
}
