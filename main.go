package main

import (
	"github.com/cherubina-gabriak/calendarApp/calendar"
	"github.com/cherubina-gabriak/calendarApp/events"
	"time"
)

func main() {
	newEvent := events.Event{
		Title:   "Новое событие",
		StartAt: time.Now(),
	}

	calendar.AddEvent("event1", newEvent)
}
