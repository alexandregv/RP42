package core

import (
	"context"
	"fmt"
	"time"

	discord "github.com/hugolgst/rich-go/client"

	"github.com/alexandregv/RP42/pkg/api"
	"github.com/alexandregv/RP42/pkg/oauth"
)

const DISCORD_APP_ID = "531103976029028367"

type presenceBody struct {
	Details        string
	State          string
	LargeText      string
	SmallImage     string
	SmallText      string
	StartTimestamp *time.Time
}

// SendActivity is the "low-level" function, which only sets the Rich Presence, with the given values.
func SendActivity(body *presenceBody) (err error) {
	err = discord.Login(DISCORD_APP_ID)
	if err != nil {
		return err
	}

	err = discord.SetActivity(discord.Activity{
		Details:    body.Details,
		State:      body.State,
		LargeImage: "logo",
		LargeText:  body.LargeText,
		SmallImage: body.SmallImage,
		SmallText:  body.SmallText,
		Timestamps: &discord.Timestamps{
			Start: body.StartTimestamp,
		},
	})
	return err
}

// BuildPresenceBody takes API values and prepare the Rich Presence body, then calls [sendActivity].
func BuildPresenceBody(ctx context.Context, user *api.User, location *api.Location, coalition *api.Coalition, campus *api.Campus) (body *presenceBody, err error) {
	cursusUser, err := user.GetPrimaryCursus()
	if cursusUser == nil {
		return nil, err
	}

	lvl := fmt.Sprintf("%.2f", cursusUser.Level)
	login := user.Login
	separator := " in "

	var (
		start      time.Time
		loc        string
		campusName string
		coaName    string
		coaSlug    string
	)

	if user.Location == "" {
		loc = "¯\\_(ツ)_/¯"
		campusName = ""
		separator = ""
		start = time.Now()
	} else {
		loc = user.Location
		start = location.BeginAt
		campusName = campus.Name
	}

	if coalition == nil {
		coaName = "None"
		coaSlug = "none"
	} else {
		coaName = coalition.Name
		coaSlug = coalition.Slug
	}

	// Discord doesn't handle Unix Epoch 0, so map it to unix0 + 1 sec
	if start.Unix() <= 0 {
		start = time.Unix(1, 0)
	}

	return &presenceBody{
		fmt.Sprintf("%s | Lvl %s", login, lvl),
		loc + separator + campusName,
		"Download: git.io/Je2xQ",
		coaSlug,
		coaName,
		&start,
	}, nil
}

// Run runs the core action of the program, calling the API to retrive info and then setting Discord Rich Presence.
func Run(ctx context.Context, login string, apiClient string, apiSecret string) (err error) {
	ctx = context.WithValue(ctx, "apiClient", oauth.NewClient(apiClient, apiSecret))

	user, err := api.GetUser(ctx, login)
	if err != nil {
		return err
	}

	loc, err := user.GetUserFirstLocation(ctx)
	if err != nil {
		return err
	}
	if loc == nil {
		time.Sleep(1 * time.Second)
		loc, err = user.GetUserLastLocation(ctx)
		if err != nil {
			return err
		}
	}

	time.Sleep(1 * time.Second)

	coa, err := user.GetUserCoalition(ctx)
	if err != nil {
		return err
	}

	campus, err := api.GetCampus(ctx, loc.CampusID)
	if err != nil {
		return err
	}

	body, err := BuildPresenceBody(ctx, user, loc, coa, campus)
	if err != nil {
		return err
	}

	return SendActivity(body)
}
