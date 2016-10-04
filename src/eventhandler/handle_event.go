package eventhandler

import (
	"event"
	"event/Docked"
	"event/DockingGranted"
	"event/DockingRequested"
	"event/EngineerProgress"
	"event/FSDJump"
	"event/FileHeader"
	"event/LoadGame"
	"event/Location"
	"event/MarketBuy"
	"event/MarketSell"
	"event/MissionAccepted"
	"event/MissionCompleted"
	"event/ModuleBuy"
	"event/ModuleSwap"
	"event/Progress"
	"event/Promotion"
	"event/Rank"
	"event/SellExplorationData"
	"event/ShipyardSwap"
	"event/SupercruiseEntry"
	"event/SupercruiseExit"
	"event/USSDrop"
	"event/Undocked"
	"fmt"
)

// struct to hold the different handlers for the json events
type HandleLine struct {
	Handlers map[string]func(*event.Event, bool) string
}

// Initalize the HandleLine struct
// by adding all of the events we want to listen for
func HandleEventInit(h *HandleLine) {
	// n stands for Name
	// f stands for function
	name, functionToCall := Docked.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = DockingGranted.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = DockingRequested.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = EngineerProgress.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = FileHeader.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = FSDJump.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = LoadGame.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = Location.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = MarketBuy.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = MarketSell.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = MissionAccepted.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = MissionCompleted.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = ModuleBuy.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = ModuleSwap.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = Progress.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = Promotion.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = Rank.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = SellExplorationData.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = ShipyardSwap.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = SupercruiseEntry.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = SupercruiseExit.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = Undocked.New()
	h.Handlers[name] = functionToCall

	name, functionToCall = USSDrop.New()
	h.Handlers[name] = functionToCall

}

func (h *HandleLine) HandleEvent(e *event.Event) {
	// if the event is valid
	if imap, ok := (*e).(map[string]interface{}); ok {
		// if it contains an event tag
		if v, ok := imap["event"]; ok {
			// if that tag is a string
			if eventStr, ok := v.(string); ok {
				// then if that tag exists in the list of valid handlers
				if fnptr, ok := h.Handlers[eventStr]; ok {
					// then print out the data that it contains
					fmt.Println(fnptr(e, false))
				}
			}
		}
	}
}
