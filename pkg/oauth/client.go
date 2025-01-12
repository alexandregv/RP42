package oauth

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	API_URL        = "https://api.intra.42.fr"
	TOKEN_ENDPOINT = "/oauth/token"
)

// Client holds an *http.Client
type Client struct {
	*http.Client
}

// NewClient creates a new [Client] using the Client Credentials OAuth flow.
func NewClient(client_id string, client_secret string) *Client {
	config := &clientcredentials.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		TokenURL:     API_URL + TOKEN_ENDPOINT,
		Scopes:       []string{},
	}

	ctx := context.Background()
	client := config.Client(ctx)

	return &Client{client}
}
