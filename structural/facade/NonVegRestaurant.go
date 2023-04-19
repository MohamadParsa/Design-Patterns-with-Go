package facade

type nonVegRestaurant struct {
	menu []nonVegFood
}
type nonVegFood struct {
	foodName string
	price    float64
}

func (nonVegRestaurant *nonVegRestaurant) getMenu() []nonVegFood {
	return nonVegRestaurant.menu
}
