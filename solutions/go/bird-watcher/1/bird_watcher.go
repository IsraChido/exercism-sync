package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var birdCount int
    
	for i:= 1; i <= len(birdsPerDay); i++ {
        birdCount += birdsPerDay[i - 1]
    }
    
    return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdCount int
    var dayToStart int
    var dayToEnd int

    if week >= 1 {
        dayToStart = (week * 7) - 6
        dayToEnd = dayToStart + 6
    } else {
        return 0
    }
    
	for i:= dayToStart; i <= dayToEnd; i++ {
        birdCount += birdsPerDay[i - 1]
    }
    
    return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:= 1; i <= len(birdsPerDay); i++ {
        if i % 2 == 1 {
            birdsPerDay[i - 1] += 1
        }
    }

    return birdsPerDay
}
