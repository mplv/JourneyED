package LoadGame

import "event"

func LoadGameEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Load game:" + sep
		if icm, ok := imap["Commander"]; ok {
			if cm, ok := icm.(string); ok {
				msg += "Commander: " + cm + sep
			}
		}
		if is, ok := imap["Ship"]; ok {
			if s, ok := is.(string); ok {
				msg += "Ship: " + s + sep
			}
		}
		if isid, ok := imap["ShipID"]; ok {
			if sid, ok := isid.(string); ok {
				msg += "Ship ID: " + sid + sep
			}
		}
		if isl, ok := imap["StartLanded"]; ok {
			if sl, ok := isl.(string); ok {
				msg += "Start landed on planet: " + sl + sep
			}
		}
		if igm, ok := imap["GameMode"]; ok {
			if gm, ok := igm.(string); ok {
				msg += "Game mode: " + gm + sep
			}
		}
		if igr, ok := imap["Group"]; ok {
			if gr, ok := igr.(string); ok {
				msg += "Group: " + gr + sep
			}
		}
		if icr, ok := imap["Credits"]; ok {
			if cr, ok := icr.(string); ok {
				msg += "Credits: " + cr + sep
			}
		}
		if il, ok := imap["Loan"]; ok {
			if l, ok := il.(string); ok {
				msg += "Loan: " + l + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "LoadGame", LoadGameEvent
}
