//Package facade is an implementation of facade pattern in Go
package facade

/*Facade is a part of Gang of Four design pattern and it is categorized under Structural design
patterns. Before we dig into the details of it, let us discuss some examples which will be solved by
this particular Pattern.
So, As the name suggests, it means the face of the building. The people walking past the road can
only see this glass face of the building. They do not know anything about it, the wiring, the pipes
and other complexities. It hides all the complexities of the building and displays a friendly face.
refrence: https://www.geeksforgeeks.org/facade-design-pattern-introduction/
*/

//GetVegMenu function returns the vegetarian food menu.
func GetVegMenu() []vegFood {
	VegRestaurant := vegRestaurant{}
	VegRestaurant.createMenu()
	return VegRestaurant.getMenu()
}

//GetNonVegMenu function returns the non vegetarian food menu.
func GetNonVegMenu() []nonVegFood {
	NonVegRestaurant := nonVegRestaurant{}
	menu := NonVegRestaurant.getMenu()
	if menu == nil {
		return []nonVegFood{{foodName: "foods are over", price: 0}}
	}
	return menu
}
