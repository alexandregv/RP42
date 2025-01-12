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

// sendActivity is the "low-level" function, which only sets the Rich Presence, with the given values.
func sendActivity(details string, state string, largeText string, smallImage string, smallText string, startTimestamp *time.Time) {
	err := discord.Login(DISCORD_APP_ID)
	if err != nil {
		panic(err)
	}

	err = discord.SetActivity(discord.Activity{
		Details:    details,
		State:      state,
		LargeImage: "logo",
		LargeText:  largeText,
		SmallImage: smallImage,
		SmallText:  smallText,
		Timestamps: &discord.Timestamps{
			Start: startTimestamp,
		},
	})
	if err != nil {
		panic(err)
	}
}

// setPresence takes API values and prepare the Rich Presence body, then calls [sendActivity].
func setPresence(ctx context.Context, user *api.User, location *api.Location, coalition *api.Coalition, campus *api.Campus) {
	cursus_user := user.GetPrimaryCursus()

	if cursus_user != nil {
		lvl := fmt.Sprintf("%.2f", cursus_user.Level)
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

		sendActivity(
			fmt.Sprintf("%s | Lvl %s", login, lvl),
			loc+separator+campusName,
			"Download: git.io/Je2xQ",
			coaSlug,
			coaName,
			&start,
		)
		return
	}
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

	setPresence(ctx, user, loc, coa, campus)
	return nil
}
