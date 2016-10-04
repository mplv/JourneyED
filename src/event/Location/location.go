package Location

import (
	"event"
	"fmt"
)

func LocationEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Current location:" + sep
		if iv, ok := imap["Docked"]; ok {
			if v, ok := iv.(float64); ok {
				if v == 1 {
					msg += "Docked at "
				}
			}
		}
		if iv, ok := imap["StationName"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Station: " + v
			}
		}

		if iv, ok := imap["StationType"]; ok {
			if v, ok := iv.(string); ok {
				msg += " (" + v + ")" + sep
			}
		}

		if iv, ok := imap["StarSystem"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Star system: " + v + sep
			}
		}
		if iv, ok := imap["StarPos"]; ok {
			if v, ok := iv.([]interface{}); ok {
				msg += "Star Pos:\n\t\tx: " + fmt.Sprint(v[0]) + ", y: " + fmt.Sprint(v[1]) + ", z: " + fmt.Sprint(v[2]) + sep
			}
		}
		if iv, ok := imap["Economy_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Economy: " + v + sep
			}
		}
		if iv, ok := imap["Government_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Government: " + v + sep
			}
		}
		if iv, ok := imap["Security_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Security: " + v + sep
			}
		}
		if iv, ok := imap["Faction"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Faction: " + v + sep
			}
		}
		if iv, ok := imap["FactionState"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Faction state: " + v + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "Location", LocationEvent
}
