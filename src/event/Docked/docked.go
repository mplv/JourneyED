package Docked

import (
	"event"
	"fmt"
)

func DockedEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {

		sep := "\n\t"
		msg := "Docked at "

		if isn, ok := imap["StationName"]; ok {
			if sn, ok := isn.(string); ok {
				msg += sn + " "
			}
		}
		if ist, ok := imap["StationType"]; ok {
			if st, ok := ist.(string); ok {
				msg += "(" + st + ") in "
			}
		}
		if iss, ok := imap["StarSystem"]; ok {
			if ss, ok := iss.(string); ok {
				msg += ss + sep
			}
		}
		if icb, ok := imap["CockpitBreached"]; ok {
			if cb, ok := icb.(bool); ok {
				msg += "Cockpit breached: " + fmt.Sprint(cb) + sep
			}
		}
		if ifa, ok := imap["Faction"]; ok {
			if fa, ok := ifa.(string); ok {
				msg += "Faction: " + fa + sep
			}
		}
		if ifs, ok := imap["FactionState"]; ok {
			if fs, ok := ifs.(string); ok {
				msg += "Faction state: " + fs + sep
			}
		}
		if ia, ok := imap["Allegiance"]; ok {
			if a, ok := ia.(string); ok {
				msg += "Allegiance: " + a + sep
			}
		}
		if iel, ok := imap["Economy_Localised"]; ok {
			if el, ok := iel.(string); ok {
				msg += "Economy: " + el + sep
			}
		}
		if igl, ok := imap["Government_Localised"]; ok {
			if gl, ok := igl.(string); ok {
				msg += "Government: " + gl + sep
			}
		}
		if isl, ok := imap["Security_Localised"]; ok {
			if sl, ok := isl.(string); ok {
				msg += "Security: " + sl + sep
			}
		}
		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "Docked", DockedEvent
}
