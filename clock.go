package main

import (
	"fmt"
	"log"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
)

// Clock implements a UI
type Clock struct {
	root                         *gtk.ApplicationWindow
	log                          *logrus.Logger
	bld                          *gtk.Builder
	analogClock                  *gtk.DrawingArea
	localTime, utcTime, unixTime *gtk.Label
}

// NewClock is a factory bean pattern constructor generator conveyor belt thing
func NewClock(log *logrus.Logger) (*Clock, error) {
	ui := &Clock{
		log: log,
	}
	err := ui.Load()
	if err == nil {

		w, err := ui.bld.GetObject("root")
		if err != nil {
			log.WithField("id", "root").WithError(err).Fatal("Cant find root widget")
		}
		ui.root = w.(*gtk.ApplicationWindow)
	}
	return ui, nil
}

// Start the clock painting
func (c *Clock) Start() {

	c.root.Connect("destroy", func() {
		gtk.MainQuit()
	})

	c.root.ShowAll()

	// make the clockArea a square of the height
	a := c.root.GetAllocation()
	height := a.Rectangle.GetHeight()
	c.analogClock.SetSizeRequest(height, height)

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

// Load sets up the widget from the gladefile
func (c *Clock) Load() error {

	b, err := gtk.BuilderNew()
	if err != nil {
		c.log.WithError(err).Fatal("Cant allocate new builder")
	}

	// find a rice.Box
	assets, err := rice.FindBox("assets")
	if err != nil {
		c.log.WithError(err).Fatal("Cant find assets box")
	}

	// get file contents as string
	gladeString, err := assets.String("clock.glade")
	if err != nil {
		c.log.WithError(err).Fatal("Cant read clock.glade")
	}

	err = b.AddFromString(gladeString)
	if err != nil {
		c.log.WithField("data", gladeString).WithError(err).Fatal("Cant build from clock.glade")
	}

	// Set the builder
	c.bld = b

	// get the clockfack widget
	w, err := b.GetObject("analogClock")
	if err != nil {
		c.log.WithField("id", "analogClock").WithError(err).Fatal("Cant find widget")
	}
	c.analogClock = w.(*gtk.DrawingArea)

	// get the 3 time labels
	w, err = b.GetObject("UTC")
	if err != nil {
		c.log.WithField("id", "UTC").WithError(err).Fatal("Cant find widget")
	}
	c.utcTime = w.(*gtk.Label)

	w, err = b.GetObject("Local")
	if err != nil {
		c.log.WithField("id", "Local").WithError(err).Fatal("Cant find widget")
	}
	c.localTime = w.(*gtk.Label)

	w, err = b.GetObject("Unixtime")
	if err != nil {
		c.log.WithField("id", "Unixtime").WithError(err).Fatal("Cant find widget")
	}
	c.unixTime = w.(*gtk.Label)

	return nil
}

// Draw does whatever is needed to re-render the UI
func (c *Clock) Draw() error {
	t := time.Now()

	tutc := t.UTC()
	tlocal := t.Local()

	c.utcTime.SetMarkup(fmt.Sprintf("<span font-family='crystal' size='x-large' foreground='#53ad95'>%02d:%02d</span>", tutc.Hour(), tutc.Minute()))
	c.localTime.SetMarkup(fmt.Sprintf("<span font-family='crystal' size='x-large' foreground='#d6c08b'>%02d:%02d</span>", tlocal.Hour(), tlocal.Minute()))
	c.unixTime.SetMarkup(fmt.Sprintf("<span font-family='crystal' size='small' foreground='#888888'>%d</span>", tutc.Unix()))

	c.analogClock.QueueDraw()

	return nil
}
