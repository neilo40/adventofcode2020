package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	//part1()
	part2()
}

// Bus describes a bus
type Bus struct {
	ID           int64
	Departure    int64
	DelayMinutes int64
}

func part1() {
	lines := common.ReadFileString("day13.input")
	earliestDeparture, _ := strconv.ParseInt(lines[0], 10, 64)
	buses := strings.Split(lines[1], ",")
	busDepartures := make([]Bus, 0)
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		var busDeparture int64 = 0
		busSchedule, _ := strconv.ParseInt(bus, 10, 64)
		for {
			busDeparture += busSchedule
			if busDeparture >= earliestDeparture {
				busDepartures = append(busDepartures, Bus{busSchedule, busDeparture, 0})
				break
			}
		}
	}
	sort.Slice(busDepartures, func(i, j int) bool {
		return busDepartures[i].Departure < busDepartures[j].Departure
	})
	waitingTime := busDepartures[0].Departure - earliestDeparture
	fmt.Printf("Result: %d\n", waitingTime*busDepartures[0].ID)
}

func part2() {
	lines := common.ReadFileString("day13.input")
	busIDs := strings.Split(lines[1], ",")
	buses := make([]Bus, 0)
	for i, busIDText := range busIDs {
		if busIDText == "x" {
			continue
		}
		busID, _ := strconv.ParseInt(busIDText, 10, 64)
		buses = append(buses, Bus{busID, 0, int64(i)})
	}

	buses = advanceTo(100000000000000, buses)
	fmt.Printf("Starting at timestamp %d\n", buses[0].Departure)

	for {
		for i, bus := range buses {
			comparisonBus := i - 1
			if i == 0 {
				comparisonBus = 1
			}
			buses[i] = iterateBus(buses[comparisonBus].Departure, bus)
		}

		if patternFound(buses) {
			fmt.Printf("Earliest timestamp: %d\n", buses[0].Departure)
			break
		}
	}
}

func advanceTo(target int64, buses []Bus) []Bus {
	for i, bus := range buses {
		buses[i].Departure = target - (target % bus.ID)
	}
	return buses
}

func iterateBus(target int64, bus Bus) Bus {
	for {
		if bus.Departure > target {
			return bus
		}
		bus.Departure += bus.ID
	}
}

func patternFound(buses []Bus) bool {
	//fmt.Printf("Current timestamp: %d\n", buses[0].Departure)
	for i := 1; i < len(buses); i++ {
		if buses[i].Departure != buses[0].Departure+buses[i].DelayMinutes {
			return false
		}
	}
	return true
}
