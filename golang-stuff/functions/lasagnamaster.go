package functions

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) (time int) {
	if minutes == 0 {
		minutes = 2
	}
	time = len(layers) * minutes
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodleCount int = 0
	var sauceCount float64 = 0.0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodleCount++
		case "sauce":
			sauceCount++
		default:
			continue
		}
	}
	return noodleCount * 50, sauceCount * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	copy(scaledQuantities, quantities)
	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] = ((scaledQuantities[i] / 2.0) * float64(portions))
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
