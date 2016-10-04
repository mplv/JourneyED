package Promotion

import (
	"event"
	"fmt"
)

func PromotionEvent(e *event.Event, squelch bool) string {
	imap, ok := (*e).(map[string]interface{})
	if ok {
		combat := 0.0
		trade := 0.0
		explore := 0.0
		empire := 0.0
		federation := 0.0
		cqc := 0.0

		if icom, ok := imap["Combat"]; ok {
			if com, ok := icom.(float64); ok {
				combat = com
			}
		}
		if itrade, ok := imap["Trade"]; ok {
			if td, ok := itrade.(float64); ok {
				trade = td
			}
		}
		if iex, ok := imap["Explore"]; ok {
			if ex, ok := iex.(float64); ok {
				explore = ex
			}
		}
		if iem, ok := imap["Empire"]; ok {
			if em, ok := iem.(float64); ok {
				empire = em
			}
		}
		if ifed, ok := imap["Federation"]; ok {
			if fed, ok := ifed.(float64); ok {
				federation = fed
			}
		}
		if icq, ok := imap["CQC"]; ok {
			if cq, ok := icq.(float64); ok {
				cqc = cq
			}
		}

		nranks := []float64{combat, trade, explore, empire, federation, cqc}
		areas := []string{"Combat", "Trade", "Explore", "Empire", "Federation", "CQC"}
		sep := "\n\t"
		msg := "Progress:" + sep
		for i := range areas {
			if nranks[i] != 0.0 && nranks[i] > 0.0 {
				msg = msg + areas[i] + ": " + fmt.Sprint(nranks[i]) + sep
			}
		}

		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "Promotion", PromotionEvent
}
