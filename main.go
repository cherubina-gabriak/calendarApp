package main

import (
	"fmt"
	"github.com/cherubina-gabriak/calendarApp/calendar"
	"github.com/cherubina-gabriak/calendarApp/events"
)

func main() {
	event1, err := events.NewEvent("Новое событие", "July 30 2025 18:00")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	calendar.AddEvent("event1", event1)

	event2, err := events.NewEvent("Haidresser appointment", "August 10 2025 15:00")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	calendar.AddEvent("event2", event2)

	calendar.ChangeEvent("event2", "Haidresser appointment", "August 10 2025 12:00")

	calendar.ShowEvents()

	calendar.DeleteEvent("event1")

	calendar.ShowEvents()
}
