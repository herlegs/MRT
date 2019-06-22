package core

// Mode is for differentiating cost of graph traversal
type Mode = string

const (
	NoCost Mode = "noCost"
	Peak   Mode = "peak"
	Night  Mode = "night"
	Normal Mode = "normal"
)

type Action = string

const (
	TakeTrain Action = "takeTrain"
	Exchange  Action = "exchange"
)

const NotReachable = -1

// Location is a place
type Location struct {
	// Name example: Jurong East
	Name string
	// Stations this location has
	Stations []*Station
}

// Station is a MRT stop
type Station struct {
	// Name example: NS1
	Name string
	// Line example: NS
	Line string
	// Other connected stations
	Neighbors []*Station
	// OpenDate example: 10 March 1990
	OpenDate string
	// Station's belonging location
	Location *Location
}

type TrafficCost struct {
	// take train per stop cost, by line and Mode
	Train map[string]map[Mode]int `json:"train"`
	// wait train, by Mode (since there is no requirement by line yet)
	Wait map[Mode]int `json:"wait"`
}
