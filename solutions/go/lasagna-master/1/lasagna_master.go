package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
    
    prepareTime := time

    if prepareTime <= 0 {
        prepareTime = 2
    }

    return len(layers) * prepareTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
		
	
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, ingredientList []string) {
	lastFriendIndex := len(friendList) - 1
	myLastIndex := len(ingredientList) - 1

	ingredientList[myLastIndex] = friendList[lastFriendIndex]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numPortion int) []float64 {
	scaled := make([]float64, len(amounts))
	scaleFactor := float64(numPortion) / 2.0

	for i, amount := range amounts {
		scaled[i] = amount * scaleFactor
	}

	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
