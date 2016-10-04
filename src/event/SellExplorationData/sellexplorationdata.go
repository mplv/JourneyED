package SellExplorationData

import (
	"event"
	"fmt"
)

func SellExplorationDataEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Sell data:" + sep
		if iv, ok := imap["Systems"]; ok {
			if v, ok := iv.([]interface{}); ok {
				msg += "Systems: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["Discovered"]; ok {
			if v, ok := iv.([]interface{}); ok {
				msg += "Discovered: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["BaseValue"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Base value: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["Bonus"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Bonus: " + fmt.Sprint(v) + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "SellExplorationData", SellExplorationDataEvent
}
