package MarketBuy

import (
	"event"
	"fmt"
)

func MarketBuyEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Market buy:" + sep

		if iv, ok := imap["Type"]; ok {
			if v, ok := iv.(string); ok {
				msg += "Item name: " + v + sep
			}
		}
		if iv, ok := imap["Count"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Count: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["BuyPrice"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Unit price: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["TotalCost"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Total cost: " + fmt.Sprint(v) + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "MarketBuy", MarketBuyEvent
}
