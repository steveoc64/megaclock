package main

//go:generate rice embed-go

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
)

func main() {
	gtk.Init(nil)

	log := logrus.New()
	clock, err := NewClock(log)
	if err != nil {
		log.Fatal("Unable to create Clock UI", err)
	}
	clock.Start()

	gtk.Main()
}
