package calendar

import (
	"fmt"
	"github.com/cherubina-gabriak/calendarApp/events"
)

var eventMap = make(map[string]events.Event)

func AddEvent(key string, event events.Event) {
	eventMap[key] = event
	fmt.Println("Событие добавлено:", event.Title)
}
