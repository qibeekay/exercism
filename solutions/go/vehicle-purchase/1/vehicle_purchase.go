package purchase

import (
    "strings"
    "fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true
    }

    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    option1Lower := strings.ToLower(option1)
    option2Lower := strings.ToLower(option2)
	result := strings.Compare(option1Lower, option2Lower) 

    if result == -1 {
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    }

    if result == 1 {
        return fmt.Sprintf("%s is clearly the better choice.", option2)
    }

    return "Any of them is fine"
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    if age < 3 {
        return originalPrice * (80.0 / 100)
    } else if age >= 10 {
        return originalPrice * (50.0 / 100)
    }
        return originalPrice * (70.0 / 100)

}
