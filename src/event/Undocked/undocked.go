package Undocked

import (
	"event"
)

func UndockedEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Undock:" + sep

		if iv, ok := imap["StationName"]; ok {
			if v, ok := iv.(string); ok {
				msg += "From: " + v
			}
		}

		if iv, ok := imap["StationType"]; ok {
			if v, ok := iv.(string); ok {
				msg += " (" + v + ")" + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "Undocked", UndockedEvent
}
