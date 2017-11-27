package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	utcLabel   *gtk.Label
	localLabel *gtk.Label
	ticksLabel *gtk.Label
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

	win.Add(createUI())

	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			_, err := glib.IdleAdd(generateLabels)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
		}
	}()

	win.ShowAll()
	gtk.Main()
}

// create the UI
func createUI() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	utcLabel, err = gtk.LabelNew("UTC 00:00:00")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	localLabel, err = gtk.LabelNew("Local 00:00:00")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	ticksLabel, err = gtk.LabelNew("Ticks 000000000")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	grid.Add(utcLabel)
	grid.Add(localLabel)
	grid.Add(ticksLabel)
	utcLabel.SetHExpand(true)
	utcLabel.SetVExpand(true)
	localLabel.SetHExpand(true)
	localLabel.SetVExpand(true)
	ticksLabel.SetHExpand(true)
	ticksLabel.SetVExpand(true)
	return &grid.Container.Widget
}

func generateLabels() bool {
	fmt.Println("Generating labels")
	return false
}
