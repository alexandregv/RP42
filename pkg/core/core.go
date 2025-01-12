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

func getActiveCursus(user *api.User) *api.CursusUser {
	var active_cursus *api.CursusUser
	for _, cursus_user := range user.CursusUsers {
		if cursus_user.Cursus.Slug == "c-piscine" && active_cursus == nil {
			active_cursus = &cursus_user
		}

		if cursus_user.Cursus.Slug == "42" && (active_cursus == nil || active_cursus.Cursus.Slug == "c-piscine") {
			active_cursus = &cursus_user
		}

		if cursus_user.Cursus.Slug == "42cursus" {
			active_cursus = &cursus_user
		}
	}
	return active_cursus
}

func setPresence(ctx context.Context, user *api.User, location *api.Location, coalition *api.Coalition, campus *api.Campus) {
	cursus_user := getActiveCursus(user)

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

func Run(ctx context.Context, login string, apiClient string, apiSecret string) (err error) {
	ctx = context.WithValue(ctx, "apiClient", oauth.NewClient(apiClient, apiSecret))

	user, err := api.GetUser(ctx, login)
	if err != nil {
		return err
	}

	loc, err := api.GetUserFirstLocation(ctx, user)
	if err != nil {
		return err
	}
	if loc == nil {
		time.Sleep(1 * time.Second)
		loc, err = api.GetUserLastLocation(ctx, user)
		if err != nil {
			return err
		}
	}

	time.Sleep(1 * time.Second)

	coa, err := api.GetUserCoalition(ctx, user)
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
