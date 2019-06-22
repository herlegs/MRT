package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetShortestPath(t *testing.T) {
	LoadGraph("../stationmap.csv", "../traffic.json")

	from, to, mode := LocationMap["hollandvillage"], LocationMap["bugis"], NoCost
	stationPath, cost := GetShortestPath(from, to, mode)
	path := make([]string, len(stationPath))
	for i, s := range stationPath {
		path[i] = s.Name
	}

	require.Equal(t, 8, cost)
	require.EqualValues(t, []string{"CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12", "DT13", "DT14"}, path)

	from, to, mode = LocationMap["boonlay"], LocationMap["littleindia"], Peak
	stationPath, cost = GetShortestPath(from, to, mode)
	path = make([]string, len(stationPath))
	for i, s := range stationPath {
		path[i] = s.Name
	}

	require.Equal(t, 150, cost)
	require.EqualValues(t, []string{"EW27", "EW26", "EW25", "EW24", "EW23", "EW22", "EW21", "CC22", "CC21", "CC20", "CC19", "DT9", "DT10", "DT11", "DT12"}, path)

}
