package SupercruiseExit

import (
	"event"
)

func SupercruiseExitEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Supercruise exit:" + sep

		if iv, ok := imap["StarSystem"]; ok {
			if v, ok := iv.(string); ok {
				msg += "In: " + v + sep
			}
		}

		if iv, ok := imap["Body"]; ok {
			if v, ok := iv.(string); ok {
				msg += "At: " + v + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "SupercruiseExit", SupercruiseExitEvent
}
