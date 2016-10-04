package FileHeader

import "event"

func FileHeaderEvent(e *event.Event, squelch bool) string {
	return "New log file"
}

func New() (string, func(*event.Event, bool) string) {
	return "fileheader", FileHeaderEvent
}
