package api

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/alexandregv/RP42/pkg/oauth"
)

const URL = "https://api.intra.42.fr"

// fetch() queries an endpoint of the API.
func fetch(ctx context.Context, endpoint string) (body []byte, err error) {
	client := ctx.Value("apiClient").(*oauth.Client)

	resp, err := client.Get(fmt.Sprint(URL, endpoint))
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return body, errors.New(fmt.Sprintf("The API responded with a bad status code (%d): %s", resp.StatusCode, string(body)))
	}

	return body, nil
}
