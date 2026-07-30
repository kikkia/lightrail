package main

import (
	"encoding/json"
	"fmt"
)

type StopTime struct {
	ArrivalEnabled   bool  "json:arrivalEnabled"
	ArrivalTime      int64 "json:arrivalTime"
	DepartureEnabled bool  "json:departureEnabled"
	DepartureTime    int64 "json:departureTime"
}

type DirectionalSchedule struct {
	ScheduleStopTimes []StopTime "json:scheduleStopTimes"
}

type RouteSchedule struct {
	RouteId                     string                "json:routeId"
	StopRouteDirectionSchedules []DirectionalSchedule "json:stopRouteDirectionSchedules"
}

type EntryStruct struct {
	Date               uint64          "json:date"
	StopId             string          "json:stopId"
	StopRouteSchedules []RouteSchedule "json:stopRouteSchedules"
}

type DataStruct struct {
	Entry EntryStruct "json:entry"
}

type Response struct {
	Code int        "json:code"
	Data DataStruct "json:data"
}

func ParseResponse(data []byte) (Response, error) {
	var response Response
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, fmt.Errorf("JSON decoding failed: %w", err)
	}
	return response, nil
}
