package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var result bool
	if kind == "car" || kind == "truck" {
		result = true
	} else {
		result = false
	}
	return result
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 + " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var result float64
	if age < 3 {
		result = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		result = originalPrice * 0.7
	} else {
		result = originalPrice * 0.5
	}
	return result
}
