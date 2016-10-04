package DockingRequested

import (
	"event"
)

func DockingRequestedEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Docking requested:" + sep

		if iv, ok := imap["StationName"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Station: " + v + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "DockingRequested", DockingRequestedEvent
}
