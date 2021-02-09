package main

import (
	"fmt"
	"os/user"
	"strings"
	"sync"
	"time"

	"github.com/alexandregv/RP42/internal/icon"
	"github.com/alexandregv/RP42/pkg/api"
	"github.com/getlantern/systray"
	discord "github.com/hugolgst/rich-go/client"
)

const DISCORD_APP_ID = "531103976029028367"

func main() {
	systray.Run(onReady, onExit)
}

func setupTray() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("RP42")

	mQuit := systray.AddMenuItem("Quit", "Quit")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
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

func setPresence(user *api.User, location *api.Location, coalition *api.Coalition) {
	cursus_user := getActiveCursus(user)

	if cursus_user != nil {
		lvl := fmt.Sprintf("%.2f", cursus_user.Level)
		login := user.Login
		campus := user.Campus[0].Name
		separator := " in "

		var (
			start   time.Time
			loc     string
			coaName string
			coaSlug string
		)

		if user.Location == "" {
			loc = "¯\\_(ツ)_/¯"
			campus = ""
			separator = ""
			start = time.Now()
		} else {
			loc = user.Location
			start = location.BeginAt
		}

		if coalition == nil {
			coaName = "None"
			coaSlug = "none"
		} else {
			coaName = coalition.Name
			coaSlug = coalition.Slug
		}

		sendActivity(
			fmt.Sprintf("%s | Lvl %s", login, lvl),
			fmt.Sprint(loc, separator, campus),
			"Download: git.io/Je2xQ",
			coaSlug,
			coaName,
			&start,
		)
		return
	}
}

func onReady() {
	setupTray()

	osUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	login := strings.ToLower(osUser.Username)

	user := api.GetUser(login)
	loc := api.GetUserFirstLocation(user)
	if loc == nil {
		time.Sleep(1 * time.Second)
		loc = api.GetUserLastLocation(user)
	}
	time.Sleep(1 * time.Second)
	coa := api.GetUserCoalition(user)

	setPresence(user, loc, coa)

	fmt.Println("Sleeping... Press CTRL+C to stop.")
	m := sync.Mutex{}
	m.Lock()
	m.Lock()
}

func onExit() {
	// clean up here
}
