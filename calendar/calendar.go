package calendar

import (
	"fmt"
	"github.com/cherubina-gabriak/calendarApp/events"
)

var eventMap = make(map[string]events.Event)

func AddEvent(key string, event events.Event) {
	eventMap[key] = event
	fmt.Println(">> Событие добавлено:", event.Title)
}

func ChangeEvent(key string, title string, startTime string) error {
	updatedEvent, err := events.NewEvent(title, startTime)

	if err != nil {
		return err
	}

	eventMap[key] = updatedEvent
	fmt.Println(">> Событие успешно изменено:", updatedEvent.Title)

	return nil
}

func DeleteEvent(key string) {
	delete(eventMap, key)
}

func ShowEvents() {
	fmt.Println(">> Список событий:")

	for _, event := range eventMap {
		fmt.Println(event.Title, '-', event.StartAt)
	}
}
