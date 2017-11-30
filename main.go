package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"

	"github.com/steveoc64/megaclock/ui"
)

func main() {
	gtk.Init(nil)

	clock, err := ui.NewClockUI()
	if err != nil {
		log.Fatal("Unable to create Clock UI", err)
	}
	w := clock.Widget().(*gtk.ApplicationWindow)
	w.Connect("destroy", func() {
		gtk.MainQuit()
	})
	w.ShowAll()

	redrawThread(clock)
	gtk.Main()
}
