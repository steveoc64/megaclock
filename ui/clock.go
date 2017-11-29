package ui

import (
	"fmt"
	"math"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// var clockUITimes = []string{
// 	"UTC", "Local", "Unix", "Nano",
// }
var clockUITimes = []string{
	"UTC", "Local", "Unix",
}

// ClockUI implements a UI
type ClockUI struct {
	grid   *gtk.Grid
	labels map[string]*gtk.Label
	times  map[string]*gtk.Label
	analog *gtk.DrawingArea
}

// NewClockUI is a factory bean pattern constructor generator conveyor belt thing
func NewClockUI() (*ClockUI, error) {
	ui := &ClockUI{}
	_, err := ui.Create()
	return ui, err
}

// Widget returns the root widget for this UI
func (c *ClockUI) Widget() gtk.IWidget {
	return c.grid
}

// Create sets up the widget heirachy of this UI
func (c *ClockUI) Create() (gtk.IWidget, error) {
	c.labels = make(map[string]*gtk.Label)
	c.times = make(map[string]*gtk.Label)

	g, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}
	c.grid = g

	g.SetOrientation(gtk.ORIENTATION_VERTICAL)
	g.SetColumnSpacing(20)
	g.SetColumnHomogeneous(true)

	row := 1
	for _, v := range clockUITimes {
		l, err := gtk.LabelNew(v)
		if err != nil {
			return nil, err
		}

		l.SetHExpand(true)
		l.SetHAlign(gtk.ALIGN_FILL)
		l.SetVExpand(true)
		// l.SetJustify(gtk.JUSTIFY_LEFT)
		c.labels[v] = l

		c.grid.Attach(l, 2, row, 1, 1)

		t, err := gtk.LabelNew("")
		if err != nil {
			return nil, err
		}
		t.SetHExpand(true)
		t.SetHAlign(gtk.ALIGN_FILL)
		t.SetVExpand(true)
		// t.SetJustify(gtk.JUSTIFY_RIGHT)
		t.SetName(v)

		c.times[v] = t
		c.grid.Attach(t, 3, row, 1, 1)

		row++
	}

	a, err := gtk.DrawingAreaNew()
	if err != nil {
		return nil, err
	}

	c.analog = a
	c.grid.Attach(a, 1, 1, 1, 3)

	// Event handlers
	a.Connect("draw", func(a *gtk.DrawingArea, cc *cairo.Context) {
		alloc := c.grid.GetAllocation()
		h := float64(alloc.GetHeight())

		cc.Scale(h*2, h*2)
		cc.Translate(h, h)
		cc.SetLineWidth(2)

		cc.Save()

		// Draw the background
		cc.SetSourceRGBA(1, 1, 1, 0.5)
		cc.Arc(1, 1, 1, 0, 2.0*math.Pi)
		cc.FillPreserve()
		cc.Restore()
		cc.Stroke()

		// Local Hour
		// t := time.Now().Local()
		// hLocal := t.Hour()
		// fmt.Println("l is", hLocal)
		// hNth := hLocal % 12
		// fmt.Println("nth is", hNth)
		// hFrac := float64(hNth) / 12.0
		// fmt.Println("each nth is a 12th of a slice of the circle", hFrac)
		// fmt.Println("start radians = ", hFrac*2.0*math.Pi)

		// mins := t.Minute()
		// fmt.Println("min is", mins)
		// minFrac := float64(mins) / 60.0
		// fmt.Println("mins is a frac of an hour", minFrac)
		// fmt.Println("end radians = ", minFrac*2.0*math.Pi)
		// ctx.SetSourceRGBA(0, 0, 1, 0.6)
		// ctx.Arc(aHeight/2, aHeight/2, aHeight/2, hFrac*2.0*math.Pi, minFrac*2.0*math.Pi)
		// ctx.Fill()

	})

	// if good, then call the first draw
	return c.grid, c.Draw()
}

// style is a pvt func to style this UI
func (c *ClockUI) style() error {
	css, err := gtk.CssProviderNew()
	if err != nil {
		return err
	}

	err = css.LoadFromData(`
GtkWindow {
background-color:grey;
    border-radius: 15px;
}
#UTC {
    background: green;
    color: white;
    font-family: DejaVu Sans;
    font-style: normal;
    font-weight: bold;
    font-size: 20px;
    border-radius: 15px;
}
 
#Local {
    background: blue;
    color: white;
    font-family: DejaVu Sans;
    font-style: normal;
    font-weight: bold;
    font-size: 20px;
    border-radius: 15px;
}
 
#Unix {
    background: red;
    color: white;
    font-family: DejaVu Sans;
    font-style: normal;
    font-weight: bold;
    font-size: 20px;
    border-radius: 15px;
}
 
#Nano {
    background: green;
    color: white;
    font-family: DejaVu Sans;
    font-style: normal;
    font-weight: bold;
    font-size: 20px;
    border-radius: 15px;
 
}
 
#UTC:hover,
#Local:hover,
#Unix:hover,
#Nano:hover {
 background-color:black;
}	
	`)

	return err
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

	c.times["Unix"].SetMarkup(fmt.Sprintf("<span font-family='crystal' size='small' foreground='#888888'>%d</span>", tutc.Unix()))
	// c.times["Nano"].SetMarkup(fmt.Sprintf("%d", t.UnixNano()))

	c.grid.QueueDraw()

	return nil
}
