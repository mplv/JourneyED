package MissionAccepted

import (
	"event"
	"fmt"
)

func MissionAcceptedEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Mission accepted:" + sep

		if iv, ok := imap["Faction"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Faction: " + v + sep
			}
		}
		if iv, ok := imap["Name"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Mission name: " + v + sep
			}
		}
		if iv, ok := imap["Commodity_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Commodity: " + v + sep
			}
		}
		if iv, ok := imap["Count"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Count: " + fmt.Sprint(v) + sep
			}
		}

		if iv, ok := imap["Expiry"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Expiry: " + v + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "MissionAccepted", MissionAcceptedEvent
}
