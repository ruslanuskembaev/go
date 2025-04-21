package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noddlesQtts int
	var sauceQtts float64
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noddlesQtts += 50
		}
		if layers[i] == "sauce" {
			sauceQtts += 0.2
		}
	}
	return noddlesQtts, sauceQtts
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] / 2 * float64(portions)
	}
	return scaledQuantities
}
