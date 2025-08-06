package lasagna

const (
    noodlesQuantity = 50
    sauceQuantity = 0.2
    stdPrepTime = 2
    stdPortion = 2
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
    if layerPrepTime <= 0 {
        layerPrepTime = stdPrepTime
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
			noodleQ += noodlesQuantity
        case "sauce":
        	sauceQ += sauceQuantity
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
        newRecipe[i] = a / float64(stdPortion) * float64(portion)
    }
    return newRecipe
}

