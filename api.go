package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const url = "https://api.pugetsound.onebusaway.org/api/where/schedule-for-stop/40_E15-T1.json?key=%s"

func fetchNextTimes() ([]int64, error) {
	token := os.Getenv("LIGHTRAIL_TOKEN")
	if token == "" {
		token = "TEST"
	}
	fullUrl := fmt.Sprintf(url, token)
	response, err := http.Get(fullUrl)
	if err != nil {
		return []int64{}, fmt.Errorf("request to %s failed: %w", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []int64{}, fmt.Errorf("request failed with status: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []int64{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var parsed Response
	parsed, err = ParseResponse(body)
	//log.Printf("%s", parsed)
	var times = make([]int64, 4)
	var nextTimes = parsed.Data.Entry.StopRouteSchedules[0].StopRouteDirectionSchedules[0].ScheduleStopTimes
	ind := 0
	for p := 0; p < len(nextTimes); p++ {
		if nextTimes[p].ArrivalTime < time.Now().UnixMilli() {
			ind++
		} else {
			break
		}
	}
	for i := 0; i < 4; i++ {
		if ind+i >= len(nextTimes) {
			break
		}
		times[i] = nextTimes[ind+i].ArrivalTime
	}
	return times, nil
}
