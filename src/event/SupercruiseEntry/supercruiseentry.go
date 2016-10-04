package SupercruiseEntry

import (
	"event"
)

func SupercruiseEntryEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Supercruise entry:" + sep

		if iv, ok := imap["StarSystem"]; ok {
			if v, ok := iv.(string); ok {
				msg += "In: " + v + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "SupercruiseEntry", SupercruiseEntryEvent
}
