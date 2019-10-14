package oauth

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)
const API_URL = "https://api.intra.42.fr"
var (
	API_CLIENT_ID string
	API_CLIENT_SECRET string
)

type Client struct {
	*http.Client
}

func GetClient() *Client {
	client := newClient(
		API_CLIENT_ID,
		API_CLIENT_SECRET,
		fmt.Sprint(API_URL, "/oauth/token"),
	)
	return client
}

func newClient(client_id string, client_secret string, token_url string) *Client {
	config := &clientcredentials.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		TokenURL:     token_url,
		Scopes:       []string{},
	}

	ctx := context.Background()
	client := config.Client(ctx)

	return &Client{client}
}
