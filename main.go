package main

import (
	"fmt"
	"github.com/alexandregv/RP42/icon"
	"github.com/ananagame/rich-go/client"
	"github.com/getlantern/systray"
	"github.com/gobs/pretty"
	"strings"
	"sync"
)

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

func sendActivity(login string, level string, coalition string, location string, logstart string) {
	//start, _ := time.Parse(time.RFC822, "08 Sep 19 16:00 UTC")

	err := client.Login("531103976029028367")
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
	})

	if err != nil {
		panic(err)
	}
}

func onReady() {
	setupTray()

	sendActivity("aguiot--", "6.75", "The Alliance", "In train", "An RFC822 string")

	//user := GetUser("aguiot--")
	//fmt.Println(user.Login)
	
	loc := GetUserLastLocation("aguiot--")
	//fmt.Println(loc)
	//fmt.Printf("%#v\n", loc)
	pretty.PrettyPrint(loc)

	fmt.Println("Sleeping... Press CTRL+C to stop.")
	m := sync.Mutex{}
	m.Lock()
	m.Lock()
}

func onExit() {
	// clean up here
}
