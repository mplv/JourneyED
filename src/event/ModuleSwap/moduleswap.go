package ModuleSwap

import (
	"event"
)

func ModuleSwapEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Swap module:" + sep

		if iv, ok := imap["FromSlot"]; ok {
			if v, ok := iv.(string); ok {
				msg += "From slot: " + v + sep
			}
		}

		if iv, ok := imap["ToSlot"]; ok {
			if v, ok := iv.(string); ok {
				msg += "To slot: " + v + sep
			}
		}

		if iv, ok := imap["FromItem_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "From item: " + v + sep
			}
		}
		if iv, ok := imap["ToItem_Localised"]; ok {
			if v, ok := iv.(string); ok {
				msg += "To item: " + v + sep
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
	return "ModuleSwap", ModuleSwapEvent
}
