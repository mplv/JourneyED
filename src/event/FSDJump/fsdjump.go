package FSDJump

import (
	"event"
	"fmt"
)

func FSDJumpEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "FSD jump:" + sep
		if iv, ok := imap["StarSystem"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Star System: " + v + sep
			}
		}
		if iv, ok := imap["StarPos"]; ok {
			if v, ok := iv.([]interface{}); ok {
				msg += "Star Pos:\n\t\tx: " + fmt.Sprint(v[0]) + ", y: " + fmt.Sprint(v[1]) + ", z: " + fmt.Sprint(v[2]) + sep
			}
		}
		if iv, ok := imap["Body"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Body: " + v + sep
			}
		}
		if iv, ok := imap["JumpDist"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Jump distance: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["FuelUsed"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Fuel used: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["FuelLevel"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Fuel level: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["BoostUsed"]; ok {
			if v, ok := iv.(bool); ok {
				msg += "Boost used: " + fmt.Sprint(v) + sep
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
		if iv, ok := imap["Allegiance"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Allegiance: " + v + sep
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

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "FSDJump", FSDJumpEvent
}
