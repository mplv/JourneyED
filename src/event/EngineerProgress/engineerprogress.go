package EngineerProgress

import "event"

func EngineerProgressEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Engineer progress:" + sep

		if iv, ok := imap["Engineer"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Engineer: " + v + sep
			}
		}
		if iv, ok := imap["Progress"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Progress: " + v + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "EngineerProgress", EngineerProgressEvent
}
