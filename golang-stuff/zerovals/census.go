package zerovals

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	var newResidentPtr *Resident
	newResidentPtr = &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return newResidentPtr
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || len(r.Address) == 0 {
		return false
	} else {
		for k, v := range r.Address {
			if v == "" || k != "street" {
				return false
			}
		}
		return true
	}
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var count int = 0
	for _, residentPtr := range residents {
		if residentPtr.HasRequiredInfo() {
			count++
		}
	}
	return count
}
