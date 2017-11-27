package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"

	"github.com/steveoc64/megaclock/ui"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	clock, err := ui.NewClockUI()
	if err != nil {
		log.Fatal("Unable to create Clock UI", err)
	}
	win.Add(clock.Widget())
	win.ShowAll()

	redrawThread(clock)
	gtk.Main()
}
