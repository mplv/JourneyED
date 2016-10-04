package event

type BasicEvent struct {
	Timestamp string
	Event     string
}

type KillerShipEntry struct {
	Name string
	Ship string
	Rank string
}

type RingEntry struct {
	Name      string
	RingClass string
	MassMT    float64
	InnerRad  float64
	OuterRad  float64
}

// sadly synthesis uses Materials as well as planets
// also annoying neither description mesh well
// so we have 2 defs for Materials one is this
// and the other is map[string]int
type MaterialEntry struct {
	Name                string
	PercentageOccurence float64
}

type IngredientEntry map[string]int

/*
type Event struct {
	BasicEvent
	Part                  int
	Language              string
	Gameversion           string
	Build                 string
	Name                  string
	Package               string
	Commander             string
	Ship                  string
	ShipID                int
	GameMode              string
	Group                 string
	Credits               int
	Loan                  int
	Combat                int
	Trade                 int
	Explore               int
	Empire                int
	Federation            int
	CQC                   int
	CockpitBreach         bool
	StartLanded           bool
	StationName           string
	StationType           string
	StarSystem            string
	Faction               string
	FactionState          string
	Allegiance            string
	Economy               string
	Economy_Localised     string
	Government            string
	Government_Localised  string
	Security              string
	Security_Localised    string
	Reason                string
	LandingPad            int
	StarPos               []float64
	Body                  string
	JumpDist              float64
	FuelUsed              float64
	FuelLevel             float64
	BoostUsed             int
	Latitude              float64
	Longitude             float64
	Docked                int
	Target                string
	Reward                int
	VictimFaction         string
	AwardingFaction       string
	KillerName            string
	KillerName_Localised  string
	KillerShip            string
	KillerRank            string
	Killers               []KillerShipEntry
	Interdictor           string
	IsPlayer              bool
	Health                float64
	Submitted             bool
	Success               bool
	Interdicted           string
	CombatRank            int
	ShieldsUp             bool
	Bodyname              string
	DistanceFromArrivalLS float64
	StarType              string
	StellarMass           float64
	Radius                float64
	AbsoluteMagnitude     float64
	OrbitalPerios         float64
	RotationPeriod        float64
	Rings                 []RingEntry
	TidalLock             int
	TerraformState        string
	PlanetClass           string
	Atmosphere            string
	Volcanism             string
	SurfaceGravity        float64
	SurfaceTemperature    float64
	SurfacePressure       float64
	Landable              bool
	Materials             []interface{} // note planets material lists and synthesis use this
	Category              string
	Count                 int
	DiscoveryNumber       int
	Cost                  int
	System                string
	Systems               []string
	Discovered            []string
	BaseValue             int
	Bonus                 int
	FileName              string
	Width                 int
	Height                int
	Type                  string
	Stolen                bool
	Abandoned             bool
	BuyPrice              int
	TotalCost             int
	SellPrice             int
	TotalSale             int
	AvgPricePaid          int
	IllegalGoods          bool
	StolenGoods           bool
	BlackMarket           bool
	Role                  string
	Engineer              string
	Blueprint             string
	Level                 int
	Override              bool
	Ingredients           []IngredientEntry
	Rank                  int
	Progress              string
	MissionID             string
	Commodity             string
	Commodity_Localised   string
	TargetType            string
	TargetFaction         string
	Expiry                string
	DestinationSystem     string
	DestinationStation    string
	PassengerCount        int
	PassengerVIPs         bool
	PassengerWanted       bool
	PassengerType         string
	Donation              int
	PermitsAwarded        []string
	Slot                  string
	BuyItem               string
	SellItem              string
	RetrievedItem         string
	EngineerModifications string
	SwapOutItem           bool
	Replacementitem       string
	FromSlot              string
	ToSlot                string
	FromItem              string
	ToItem                string
	Amount                float64
	BrokerPercentage      float64
	Item                  string
	Loadout               string
	ShipPrice             int
	StoreOldShip          string
	StoreShipID           int
	SellOldShip           string
	SellShipID            int
	Distance              float64
	TransferPrice         int
	Power                 string
	FromPower             string
	ToPower               string
	Votes                 int
	CrimeType             string
	Fine                  int
	Bounty                int
	Message               string
	PayeeFaction          string
	Scooped               float64
	Total                 float64
	BoostValue            float64
	Module                string
	PlayerControlled      bool
	Modules               []string
	From                  string
	Option                string
	Bankrupt              bool
	To                    string
	USSType               string
	USSThreat             int
	Others                []string
}
*/
