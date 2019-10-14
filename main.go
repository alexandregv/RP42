package main

import (
	"fmt"
	"github.com/alexandregv/RP42/icon"
	"github.com/ananagame/rich-go/client"
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
	err := client.Login(DISCORD_APP_ID)
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(client.Activity{
		Details:    fmt.Sprintf("Level: %s", level),
		State:      fmt.Sprintf("Location: %s", location),
		LargeImage: "logo",
		LargeText:  login,
		SmallImage: strings.ToLower(strings.Replace(coalition, " ", "-", -1)),
		SmallText:  coalition,
		Timestamps: &client.Timestamps{
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

	user := GetUser(login)
	loc := GetUserLastLocation(login)
	time.Sleep(1 * time.Second)
	coa := GetUserCoalition(login)

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
