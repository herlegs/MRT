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

//// Line is a MRT line
//type Line struct {
//	// map of station on the line and their index
//	StationIndexMap map[*Station]int
//	// all the exchange stations on the line
//	ExchangeStations map[*Station]bool
//}
//
//// GetCost of two station if and only if they are on same line
//func (l *Line) GetCost(from, to *Station, cost *TrafficCost, mode Mode) int {
//	fromIdx, fromExist := l.StationIndexMap[from]
//	toIdx, toExist := l.StationIndexMap[to]
//	if !fromExist || !toExist {
//		return NotReachable
//	}
//	return Abs(fromIdx-toIdx) * cost.Train[from.Line][mode]
//}
//
//type TransportGraph struct {
//	// map of all lines by line name
//	lines map[string]*Line
//}
