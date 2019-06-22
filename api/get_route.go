package api

import (
	"fmt"
	"time"

	"github.com/herlegs/MRT/core"
	"github.com/herlegs/MRT/pb"
	"golang.org/x/net/context"
)

const (
	NotValidLocation     = "the location you specified (from[%v]:%v, to[%v]:%v) is not valid"
	NotReachableLocation = "cannot reach %v from %v"
	SummaryTmpl          = "Travel from %v to %v"
	TimeBucketTmpl       = " during %v hours"
	TimeSpentTmpl        = "%v minutes"
)

func (s *Server) GetRoute(ctx context.Context, req *pb.RouteRequest) (*pb.RouteResponse, error) {
	response := &pb.RouteResponse{}
	timeStart := time.Now()
	defer func() {
		response.QueryTime = fmt.Sprintf("time used: %v", time.Now().Sub(timeStart))
	}()

	mode := getTrafficMode(req.Time)

	if cached, exist := s.getCache(req.From, req.To, mode); exist {
		response = cached
		return response, nil
	}

	// validate from and to location
	fromLocation, fromExist := core.GetLocation(req.From)
	toLocation, toExist := core.GetLocation(req.To)

	if !fromExist || !toExist {
		response.Summary = fmt.Sprintf(NotValidLocation, req.From, fromExist, req.To, toExist)

		return response, nil
	}

	stations, cost := core.GetShortestPath(fromLocation, toLocation, mode)
	if len(stations) == 0 || cost == core.NotReachable {
		response.Summary = fmt.Sprintf(NotReachableLocation, toLocation.Name, fromLocation.Name)
		s.setCache(req.From, req.To, mode, response)
		return response, nil
	}

	route := make([]string, len(stations))
	for i, s := range stations {
		route[i] = s.Name
	}

	response.Summary = fmt.Sprintf(SummaryTmpl, fromLocation.Name, toLocation.Name)
	response.StationTravelled = int64(len(stations)) - 1
	response.Route = route
	if mode != core.NoCost {
		response.Summary += fmt.Sprintf(TimeBucketTmpl, mode)
		response.TravelTime = fmt.Sprintf(TimeSpentTmpl, cost)
	}
	response.Instruction = formatInstruction(stations)

	s.setCache(req.From, req.To, mode, response)
	return response, nil
}

func getTrafficMode(timeStr string) core.Mode {
	if timeStr == "" {
		return core.NoCost
	}
	t, err := parseTime(timeStr)
	if err != nil {
		return core.NoCost
	}
	if isPeakHour(t) {
		return core.Peak
	} else if isNightHour(t) {
		return core.Night
	}
	return core.Normal
}

func parseTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", timeStr)
}

func isPeakHour(t time.Time) bool {
	// Peak hours (6am-9am and 6pm-9pm on Mon-Fri)
	weekday := t.Weekday()
	hour, _, _ := t.Clock()
	if weekday >= time.Monday && weekday <= time.Friday &&
		((hour >= 6 && hour <= 9) || (hour >= 18 && hour <= 21)) {
		return true
	}
	return false
}

func isNightHour(t time.Time) bool {
	// Night hours (10pm-6am on Mon-Sun)
	hour, _, _ := t.Clock()
	if (hour >= 22 && hour <= 23) || (hour >= 0 && hour <= 6) {
		return true
	}
	return false
}

func formatInstruction(stations []*core.Station) string {
	instruction := ""
	n := len(stations)
	if n == 0 {
		return instruction
	}
	instruction += fmt.Sprintf("Start at %v\n", stations[0].Location.Name)

	for i := 1; i < n; i++ {
		if stations[i].Line == stations[i-1].Line {
			instruction += fmt.Sprintf("Take %v line from %v to %v\n", stations[i].Line, stations[i-1].Location.Name, stations[i].Location.Name)
		} else {
			instruction += fmt.Sprintf("Change from %v line to %v line\n", stations[i-1].Line, stations[i].Line)
		}
	}

	instruction += fmt.Sprintf("And you have arrived %v\n", stations[n-1].Location.Name)
	return instruction
}
