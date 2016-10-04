package DockingGranted

import (
	"event"
	"fmt"
)

func DockingGrantedEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Docking granted:" + sep

		if iv, ok := imap["StationName"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Station: " + v + sep
			}
		}

		if iv, ok := imap["LandingPad"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Pad: " + fmt.Sprint(v) + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "DockingGranted", DockingGrantedEvent
}
