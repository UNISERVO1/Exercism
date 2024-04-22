package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	choice := " is clearly the better choice."
	if option1 < option2 {
		choice = option1 + choice
	} else {
		choice = option2 + choice
	}
	return choice
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	price := originalPrice

	if int(age) < 3 {
		price = 0.8 * price
	} else if int(age) >= 10 {
		price = 0.5 * price
	} else {
		price = .7 * price
	}
	return price
}
