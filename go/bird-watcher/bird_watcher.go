package birdwatcher

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	firstDayOfWeek := 7 * (week - 1)
	firstDayOfNextWeek := min(firstDayOfWeek+7, len(birdsPerDay))
	return TotalBirdCount(birdsPerDay[firstDayOfWeek:firstDayOfNextWeek])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		birdsPerDay[i] += 1 - i%2
	}
	return birdsPerDay
}
