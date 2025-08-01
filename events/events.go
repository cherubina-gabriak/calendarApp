package events

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	titleValid := IsTitleValid(title)

	if !titleValid {
		return Event{}, errors.New("Название должно быть не меньше 3 и не больше 50 знаков и состоять только из букв и цифр")
	}

	t, err := dateparse.ParseAny(dateStr)

	if err != nil {
		return Event{}, err
	}

	return Event{
		Title:   title,
		StartAt: t,
	}, nil
}
