package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// does unit exist
	unit_size, exist := units[unit]

	if exist {
		// if unit in bill already get quantity, else 0
		bill_existing_units := bill[item]
		// update quantity
		bill[item] = bill_existing_units + unit_size
	}

	return exist
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// does unit exist
	unit_size, exist := units[unit]

	if exist {
		// if unit in bill already get quantity
		bill_existing_units, exist_in_bill := bill[item]
		exist = exist_in_bill && bill_existing_units >= unit_size

		if exist {
			// update quantity
			bill[item] = bill_existing_units - unit_size

			// remove item if 0 quantity
			if bill[item] == 0 {
				delete(bill, item)
			}
		}
	}

	return exist
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
