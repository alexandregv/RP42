package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
	"sync"
	"time"

	"github.com/alexandregv/RP42/pkg/api"
	"github.com/alexandregv/RP42/pkg/oauth"
	discord "github.com/hugolgst/rich-go/client"
)

const DISCORD_APP_ID = "531103976029028367"

func main() {
	onReady()
}

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
			fmt.Sprint(loc, separator, campusName),
			"Download: git.io/Je2xQ",
			coaSlug,
			coaName,
			&start,
		)
		return
	}
}

func onReady() {
	ctx := context.Background()

	osUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	login := strings.ToLower(osUser.Username)

	var apiClient string
	var apiSecret string
	flag.StringVar(&apiClient, "i", "", "Client ID from API settings")
	flag.StringVar(&apiClient, "id", "", "Client ID from API settings")
	flag.StringVar(&apiSecret, "s", "", "Client Secret from API settings")
	flag.StringVar(&apiSecret, "secret", "", "Client Secret from API settings")
	flag.Usage = func() {
		fmt.Print(`Usage of RP42:
	-i, --id Client ID of your API app (required)
	-s, --secret Client Secret of your API app (required)
	
If you don't have an API app yet, create one here: https://profile.intra.42.fr/oauth/applications/new
/!\ Do NOT share your credentials to someone else, or on GitHub, etc. /!\

`)
	}
	flag.Parse()

	if apiClient == "" || apiSecret == "" {
		fmt.Println("Please provide Intra API credentials with --id and --secret. See --help for help.")
		os.Exit(2)
	}

	ctx = context.WithValue(ctx, "apiClient", oauth.NewClient(apiClient, apiSecret))

	user := api.GetUser(ctx, login)
	loc := api.GetUserFirstLocation(ctx, user)
	if loc == nil {
		time.Sleep(1 * time.Second)
		loc = api.GetUserLastLocation(ctx, user)
	}
	time.Sleep(1 * time.Second)
	coa := api.GetUserCoalition(ctx, user)
	campus := api.GetCampus(ctx, loc.CampusID)

	setPresence(ctx, user, loc, coa, campus)

	fmt.Println("Sleeping... Press CTRL+C to stop.")
	m := sync.Mutex{}
	m.Lock()
	m.Lock()
}
