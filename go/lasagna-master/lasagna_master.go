package lasagna

// TODO: define the 'PreparationTime()' function

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

func PreparationTime(layers []string, avgTime int) int {
	if avgTime > 0 {
		return avgTime * len(layers)
	} else {
		return 2 * len(layers)
	}
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles, sauce = 0, 0.0
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfDiners int) []float64 {
	scaledRecipe := make([]float64, len(quantities))
	for i, ingredient := range quantities {
		scaledRecipe[i] = ingredient * float64(numberOfDiners) / 2.0
	}
	return scaledRecipe
}
