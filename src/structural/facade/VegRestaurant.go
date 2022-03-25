package facade

type vegRestaurant struct {
	menu []vegFood
}
type vegFood struct {
	name  string
	price float64
}

func (vegRestaurant *vegRestaurant) getMenu() []vegFood {
	return vegRestaurant.menu
}
func (vegRestaurant *vegRestaurant) createMenu() {
	vegRestaurant.menu = []vegFood{{name: "tomato salad", price: 10}, {name: "broccoli salad", price: 10}}
}
