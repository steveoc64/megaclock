package main

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/steveoc64/megaclock/ui"
)

// spawns a goroutine to redraw all the things
func redrawThread(c *ui.ClockUI) {

	go func() {
		for {
			time.Sleep(time.Second * 1)
			_, err := glib.IdleAdd(c.Draw)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
		}
	}()
}
