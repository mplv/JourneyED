package ModuleBuy

import (
	"event"
	"fmt"
)

func ModuleBuyEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Buy module:" + sep

		if iv, ok := imap["Slot"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Slot: " + v + sep
			}
		}

		if iv, ok := imap["BuyItem_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Item bought: " + v + sep
			}
		}

		if iv, ok := imap["BuyPrice"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Item cost: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["Ship"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Ship: " + v + sep
			}
		}


		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "ModuleBuy", ModuleBuyEvent
}
