package Rank

import "event"

func RankEvent(e *event.Event, squelch bool) string {
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
		ranks := [][]string{{"Harmless", "Mostly Harmless", "Novice", "Competent", "Expert", "Master", "Dangerous", "Deadly", "Elite"},
			{"Penniless", "Mostly Penniless", "Peddler", "Dealer", "Merchant", "Broker", "Entrepreneur", "Tycoon", "Elite"},
			{"Aimless", "Mostly Aimless", "Scout", "Surveyor", "Explorer", "Pathfinder", "Ranger", "Pioneer", "Elite"},
			{"None", "Outsider", "Serf", "Master", "Squire", "Knight", "Lord", "Baron", "Viscount", "Count", "Earl", "Marquis", "Duke", "Prince", "King"},
			{"None", "Recruit", "Cadet", "Midshipman", "Petty Officer", "Chief Petty Officer", "Warrant Officer", "Ensign", "Lieutenant", "Lt. Commander", "Post Commander", "Post Captain", "Rear Admiral", "Vice Admiral", "Admiral"},
			{"Helpless", "Mostly Helpless", "Amateur", "Semi Professional", "Professional", "Champion", "Hero", "Legend", "Elite"}}
		areas := []string{"Combat", "Trade", "Explore", "Empire", "Federation", "CQC"}
		sep := "\n\t"
		msg := "Ranks:" + sep
		for i := range areas {
			msg = msg + areas[i] + ": " + ranks[i][int(nranks[i])] + sep
		}
		return msg
	}
	return ""
}

func New() (string, func(*event.Event, bool) string) {
	return "Rank", RankEvent
}
