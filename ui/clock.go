package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gotk3/gotk3/gtk"
)

// var clockUITimes = []string{
// 	"UTC", "Local", "Unix", "Nano",
// }
var clockUITimes = []string{
	"UTC", "Local", "Unixtime",
}

// ClockUI implements a UI
type ClockUI struct {
	b      *gtk.Builder
	grid   *gtk.Grid
	labels map[string]*gtk.Label
	times  map[string]*gtk.Label
	analog *gtk.DrawingArea
}

// NewClockUI is a factory bean pattern constructor generator conveyor belt thing
func NewClockUI() (*ClockUI, error) {
	ui := &ClockUI{}
	err := ui.Create()
	return ui, err
}

// Widget returns the root widget for this UI
func (c *ClockUI) Widget() gtk.IWidget {
	if c.b == nil {
		return nil
	}
	w, err := c.b.GetObject("root")
	if err != nil {
		return nil
	}
	spew.Dump("got w", w)
	return w.(gtk.IWidget)
}

// Create sets up the widget from the gladefile
func (c *ClockUI) Create() error {

	c.times = make(map[string]*gtk.Label)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Cainte builder new", err)
	}
	err = b.AddFromFile("ui/megaclock.glade")
	if err != nil {
		log.Fatal("Cant open the glade im afraid", err)
	}

	c.b = b
	for _, v := range clockUITimes {
		l, err := b.GetObject(v)
		if err != nil {
			log.Fatalf("Cant find label %s: %s", v, err)
		}
		c.times[v] = l.(*gtk.Label)
	}
	w, err := b.GetObject("drawZone")
	if err != nil {
		log.Fatal("Cant find drawZone", err)
	}
	c.analog = w.(*gtk.DrawingArea)

	return nil
}

// Draw does whatever is needed to re-render the UI
func (c *ClockUI) Draw() error {
	t := time.Now()

	tutc := t.UTC()
	// c.times["UTC"].SetMarkup(fmt.Sprintf("<span foreground='#53ad95'>%02d:%02d:%02d</span>", t1.Hour(), t1.Minute(), t1.Second()))
	c.times["UTC"].SetMarkup(fmt.Sprintf("<span font-family='crystal' size='x-large' foreground='#53ad95'>%02d:%02d</span>", tutc.Hour(), tutc.Minute()))

	tlocal := t.Local()
	c.times["Local"].SetMarkup(fmt.Sprintf("<span font-family='crystal' size='x-large' foreground='#d6c08b'>%02d:%02d</span>", tlocal.Hour(), tlocal.Minute()))
	// c.times["Local"].SetMarkup(fmt.Sprintf("<span foreground='#d6c08b'>%02d:%02d:%02d</span>", t1.Hour(), t1.Minute(), t1.Second()))
	// c.times["Unix"].SetMarkup(fmt.Sprintf("<span font-family='crystal' size='large' foreground='#888888'>%d</span>", tutc.Unix()))

	c.times["Unixtime"].SetMarkup(fmt.Sprintf("<span font-family='crystal' size='small' foreground='#888888'>%d</span>", tutc.Unix()))
	// c.times["Nano"].SetMarkup(fmt.Sprintf("%d", t.UnixNano()))

	c.analog.QueueDraw()

	return nil
}
