package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    return map[string]int {
        "quarter_of_a_dozen": 3,
		"half_of_a_dozen":	6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, ok := units[unit]
    if !ok {
        return ok
    }
	bill[item] += u
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    itemQt, ok := bill[item]
    if !ok {
        return ok
    }

	unitValue, ok := units[unit]
    if !ok {
        return ok
    }

    newQt := (itemQt - unitValue)
    switch {
    case newQt < 0:
    	return false
    case newQt == 0:
    	delete(bill, item)
    case newQt > 0:
    	bill[item] = newQt
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQt, ok := bill[item]
    return itemQt, ok
}
