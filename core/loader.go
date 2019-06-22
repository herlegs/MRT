package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
)

var (
	loadOnce sync.Once
	// map of location by location name
	LocationMap map[string]*Location

	trafficCost *TrafficCost

	stationNameReg = regexp.MustCompile("([a-zA-z]+)[0-9]+")
)

func readFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return ""
	}
	return string(b)
}

func LoadGraph(mrtFile, trafficFile string) {
	loadOnce.Do(func() {
		loadStationMap(mrtFile)
		loadTrafficCost(trafficFile)
	})
}

func loadStationMap(mrtFile string) {
	str := readFile(mrtFile)
	records := strings.Split(str, "\n")

	LocationMap = make(map[string]*Location)
	var prevStation *Station
	for _, record := range records {
		parts := strings.Split(record, "\t")
		stationName, locationName, openDate := parts[0], parts[1], parts[2]
		location, exist := LocationMap[NormalizeName(locationName)]
		if !exist {
			location = &Location{
				Name: locationName,
			}
			LocationMap[NormalizeName(locationName)] = location
		}

		lineName := parseLineName(stationName)
		station := &Station{
			Name:     stationName,
			Line:     lineName,
			OpenDate: openDate,
			Location: location,
		}

		location.Stations = append(location.Stations, station)

		if prevStation != nil && prevStation.Line == station.Line {
			prevStation.Neighbors = append(prevStation.Neighbors, station)
			station.Neighbors = append(station.Neighbors, prevStation)
		}

		prevStation = station
	}

	// connect exchange stations
	for _, location := range LocationMap {
		n := len(location.Stations)
		if n > 1 {
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					a, b := location.Stations[i], location.Stations[j]
					a.Neighbors = append(a.Neighbors, b)
					b.Neighbors = append(b.Neighbors, a)
				}
			}
		}
	}
}

func loadTrafficCost(trafficFile string) {
	str := readFile(trafficFile)
	trafficCost = &TrafficCost{}
	_ = json.Unmarshal([]byte(str), trafficCost)
	trafficCost.Wait[NoCost] = 1
	for _, cost := range trafficCost.Train {
		cost[NoCost] = 1
	}
}

func parseLineName(stationName string) string {
	return stationNameReg.FindAllStringSubmatch(stationName, -1)[0][1]
}

func NormalizeName(name string) string {
	name = strings.Replace(name, " ", "", -1)
	return strings.ToLower(name)
}

func GetLocation(locationName string) (*Location, bool) {
	l, exist := LocationMap[NormalizeName(locationName)]
	return l, exist
}
