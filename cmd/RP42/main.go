package main

import (
	"fmt"
	"github.com/alexandregv/RP42/internal/icon"
	"github.com/alexandregv/RP42/pkg/api"
	discord "github.com/ananagame/rich-go/client"
	"github.com/getlantern/systray"
	"os/user"
	"sync"
	"time"
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

func sendActivity(details string, state string, largeText string, smallImage string, smallText string, startTimestamp int64) {
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

func setPresence(user *api.User, location *api.Location, coalition *api.Coalition) {
	for _, cursus_user := range user.CursusUsers {
		if cursus_user.EndAt == nil {
			lvl := fmt.Sprintf("%.2f", cursus_user.Level)
			login := user.Login
			campus := user.Campus[0].Name
			separator := " in "

			var (
				start   int64
				loc     string
				coaName string
				coaSlug string
			)

			if user.Location == "" {
				loc = "¯\\_(ツ)_/¯"
				campus = ""
				separator = ""
				start = time.Now().Unix()
			} else {
				loc = user.Location
				start = location.BeginAt.Unix()
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
				start,
			)
			return
		}
	}

}

func onReady() {
	setupTray()

	osUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	login := osUser.Username

	user := api.GetUser(login)
	loc := api.GetUserFirstLocation(user)
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
