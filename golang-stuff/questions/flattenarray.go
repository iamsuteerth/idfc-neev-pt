package questions

func Flatten(nested interface{}) []interface{} {
	arr := []interface{}{}
	// Take advantage of type assertions
	// If the element is of slice type, open it up and add the elemenents
	switch i := nested.(type) {
	// It is a 1D thing
	case []interface{}:
		for _, val := range i {
			arr = append(arr, Flatten(val)...)
		}
	case interface{}:
		arr = append(arr, i)
	}
	return arr
}
