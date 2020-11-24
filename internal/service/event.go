package service

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type EventType int

type Event struct {
	Time    time.Time
	Type    EventType
	Message string
}

const (
	Ready = iota
	Message
	Error
	Close
)

func (Event) Create(eventType EventType, message string) Event {
	result := Event{
		Time: time.Now(),
		Type: eventType,
	}

	timeString := color.HiBlueString(result.Time.Format("[15:04:05]"))
	var eventString string

	switch eventType {
	case Ready:
		eventString = color.HiGreenString("[Ready]")
	case Message:
		eventString = color.HiGreenString("[Message]")
	case Error:
		eventString = color.HiRedString("[Error]")
	case Close:
		eventString = color.HiRedString("[Close]")
	}

	result.Message = fmt.Sprintf("%s %s: %s", timeString, eventString, message)

	return result
}
