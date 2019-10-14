package oauth

import (
	"fmt"
	"golang.org/x/net/context"
	clientcreds "golang.org/x/oauth2/clientcredentials"
	"net/http"
	"os"
)

const URL = "https://api.intra.42.fr"

type Client struct {
	*http.Client
}

func GetClient() *Client {
	client := newClient(
		os.Getenv("RP42_CLIENT_ID"),
		os.Getenv("RP42_CLIENT_SECRET"),
		fmt.Sprint(URL, "/oauth/token"),
	)
	return client
}

func newClient(client_id string, client_secret string, token_url string) *Client {
	config := &clientcreds.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		TokenURL:     token_url,
		Scopes:       []string{},
	}

	ctx := context.Background()
	client := config.Client(ctx)

	return &Client{client}
}
