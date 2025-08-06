package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
    if layerPrepTime <= 0 {
        layerPrepTime = 2
    }

    return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodleQ int
    var sauceQ float64

    for _, l := range layers {
        switch l{
        case "noodles":
			noodleQ += 50
        case "sauce":
        	sauceQ += 0.2
        }
    }

    return noodleQ, sauceQ
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredients, myIngredients []string) {
    myIngredients[len(myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1] 
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portion int) []float64{
    newRecipe := make([]float64, len(amounts))
    for i, a := range amounts {
        newRecipe[i] = a / 2.0 * float64(portion)
    }
    return newRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
