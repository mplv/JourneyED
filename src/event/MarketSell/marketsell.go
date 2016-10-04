package MarketSell

import (
	"event"
	"fmt"
)

func MarketSellEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		sep := "\n\t"
		msg := "Market sell:" + sep

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
		if iv, ok := imap["SellPrice"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Unit sell price: " + fmt.Sprint(v) + sep
			}
		}
		if iv, ok := imap["TotalSale"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Total sale: " + fmt.Sprint(v) + sep
			}
		}

		if iv, ok := imap["AvgPricePaid"]; ok {
			if v, ok := iv.(float64); ok {
				msg += "Average unit cost: " + fmt.Sprint(v) + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "MarketSell", MarketSellEvent
}
