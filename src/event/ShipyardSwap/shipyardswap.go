package ShipyardSwap

import (
	"event"
	"fmt"
)

func ShipyardSwapEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Swap ship:" + sep

		if iv, ok := imap["ShipType"]; ok {
			if v, ok := iv.(string); ok {
				msg += "To: " + v + sep
			}
		}

		if iv, ok := imap["ShipID"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "\tShip ID: " + fmt.Sprint(v) + sep
			}
		}

		if iv, ok := imap["StoreOldShip"]; ok {
			if v, ok := iv.(string); ok {
				msg += "From: " + v + sep
			}
		}
		if iv, ok := imap["StoreShipID"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "\tShip ID: " + fmt.Sprint(v) + sep
			}
		}


		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "ShipyardSwap", ShipyardSwapEvent
}
