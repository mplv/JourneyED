package USSDrop

import (
	"event"
	"fmt"
)

func USSDropEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "USS drop:" + sep

		if iv, ok := imap["USSType_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "USS type: " + v + sep
			}
		}
		if iv, ok := imap["USSThreat"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "USS threat level: " + fmt.Sprint(v) + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "USSDrop", USSDropEvent
}
