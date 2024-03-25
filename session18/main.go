package main

import (
	"fmt"
	"sort"
)

type activity struct {
	start, finish int
}

func activitySelection(activities []activity) []activity {
	// Sort activities by their finish time
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].finish < activities[j].finish
	})

	var selectedActivities []activity

	// Select the first activity
	selectedActivities = append(selectedActivities, activities[0])

	for i := 1; i < len(activities); i++ {
		// If the current activity's start time is after the finish time of the last selected activity,
		if activities[i].start >= selectedActivities[len(selectedActivities)-1].finish {
			selectedActivities = append(selectedActivities, activities[i])
		}
	}

	return selectedActivities
}

func main() {
	activities := []activity{
		{1, 2},
		{3, 4},
		{0, 6},
		{5, 7},
		{8, 9},
		{5, 9},
	}

	selectedActivities := activitySelection(activities)

	fmt.Println("Selected Activities:")
	for _, act := range selectedActivities {
		fmt.Printf("[%v, %v]\n", act.start, act.finish)
	}
}
