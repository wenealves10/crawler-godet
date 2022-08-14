package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/raff/godet"
)

func main() {
	chromeapp := "chromium-browser"
	chromeappArg := []string{"--headless", "--no-sandbox", "--hide-scrollbars", "--remote-debugging-port=9222", "--disable-gpu", "--allow-insecure-localhost"}
	cmd := exec.Command(chromeapp, chromeappArg...)
	err := cmd.Start()
	if err != nil {
		log.Println("cannot start browser", err)
	}

	// Will wait for chromium to start
	time.Sleep(5 * time.Second)

	// connect to Chromium instance
	remote, err := godet.Connect("localhost:9222", true)
	if err != nil {
		log.Println("cannot connect to Chrome instance:", err)
		return
	}

	// disconnect when done
	defer remote.Close()

	remote.PageEvents(true)
	remote.DOMEvents(true)

	_, err = remote.Navigate("https://www.google.com")
	if err != nil {
		log.Println("cannot connect to Chrome instance:", err)
		return
	}

	_ = remote.SaveScreenshot("screenshot.png", 0644, 0, true)

	time.Sleep(30 * time.Second)

	killapp := "kill"
	killappArg := []string{"$(lsof -t -i:9222)"}
	cmd = exec.Command(killapp, killappArg...)
	err = cmd.Start()
	if err != nil {
		log.Println("cannot kill processes", err)
	}
}
