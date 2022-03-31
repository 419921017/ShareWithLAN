package main

import (
	"ShareWithLAN/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/zserge/lorca"
)

func main() {

	go server.Run()
	ui := startBrowser()
	chSignal := listenToSingleChannel()

	select {
	case <-ui.Done():
	case <-chSignal:
	}

	defer func(ui lorca.UI) {
		err := ui.Close()
		if err != nil {
		}
	}(ui)

}

func startBrowser() lorca.UI {
	ui, _ := lorca.New("http://127.0.0.1:8080/static/index.html",
		"", 800, 600, "--disable-sync", "--disable-translate")
	return ui
}

func listenToSingleChannel() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	return chSignal
}
