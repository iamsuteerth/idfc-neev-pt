package maps

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMap := make(map[string]int)
	unitMap["quarter_of_a_dozen"] = 3
	unitMap["half_of_a_dozen"] = 6
	unitMap["dozen"] = 12
	unitMap["small_gross"] = 120
	unitMap["gross"] = 144
	unitMap["great_gross"] = 1728
	return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	billMap := make(map[string]int)
	return billMap
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	if !unitExists {
		return false
	}
	_, itemExists := bill[item]
	if itemExists {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	if !itemExists {
		return false
	}
	_, unitExists := units[unit]
	if !unitExists {
		return false
	}
	bill[item] -= units[unit]
	if bill[item] < 0 {
		// Reverting the changes if the operation resulted in negative quantity.
		bill[item] += units[unit]
		return false
	}
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}
	return bill[item], true
}
