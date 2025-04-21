package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var result int
	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var result int
	for i := 0 + (week-1)*7; i < (week-1)*7+7; i++ {
		result += birdsPerDay[i]

	}
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.

func FixBirdCountLog(birdsPerDay []int) []int {
	bpd := birdsPerDay
	for i := 0; i < len(birdsPerDay); i += 2 {
		bpd[i] = birdsPerDay[i] + 1
	}
	return bpd
}
