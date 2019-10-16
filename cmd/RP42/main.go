package main

import (
	"fmt"
	"github.com/alexandregv/RP42/internal/icon"
	"github.com/alexandregv/RP42/pkg/api"
	discord "github.com/ananagame/rich-go/client"
	"github.com/getlantern/systray"
	"os/user"
	"strings"
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

	mQuit := systray.AddMenuItem("Quitter", "Quitter")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func sendActivity(login string, level string, coalition string, location string, begin int64) {
	err := discord.Login(DISCORD_APP_ID)
	if err != nil {
		panic(err)
	}

	err = discord.SetActivity(discord.Activity{
		Details:    fmt.Sprintf("Level: %s", level),
		State:      fmt.Sprintf("Location: %s", location),
		LargeImage: "logo",
		LargeText:  login,
		SmallImage: strings.ToLower(strings.Replace(coalition, " ", "-", -1)),
		SmallText:  coalition,
		Timestamps: &discord.Timestamps{
			Start: begin,
		},
	})

	if err != nil {
		panic(err)
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
	loc := api.GetUserLastLocation(login)
	time.Sleep(1 * time.Second)
	coa := api.GetUserCoalition(login)

	lvl := fmt.Sprintf("%.2f", user.CursusUsers[0].Level)
	begin := loc.BeginAt.Unix()

	sendActivity(login, lvl, coa.Name, loc.Host, begin)

	fmt.Println("Sleeping... Press CTRL+C to stop.")
	m := sync.Mutex{}
	m.Lock()
	m.Lock()
}

func onExit() {
	// clean up here
}
