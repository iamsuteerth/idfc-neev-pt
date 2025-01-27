package numbers

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate * 0.01
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var productionRatePerMinute float64 = float64(productionRate) / 60.0
	return int(productionRatePerMinute * successRate * 0.01)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTen int = carsCount / 10
	var individualCars int = carsCount % 10
	return uint((95000 * groupsOfTen) + (10000 * individualCars))
}
