package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}

    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728

    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    _, unitExists := units[unit]

    if unitExists == false {
        return false
    }

    bill[item] += units[unit]

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
    _, unitExists := units[unit]

    itemQuantity := bill[item] - units[unit]

    if itemExists == false || unitExists == false || itemQuantity < 0 {
        return false
    } else if itemQuantity == 0 {
        delete(bill, item)
    } else {
        bill[item] -= units[unit]
    }

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billValue, itemExists := bill[item]

    if itemExists == false {
        return 0, false
    } else {
        return billValue, true
    }
}
