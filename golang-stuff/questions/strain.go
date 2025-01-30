package questions

func Keep[AnyCollection any](arr []AnyCollection, predicate func(AnyCollection) bool) []AnyCollection {
	var res []AnyCollection
	for _, v := range arr {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

func Discard[AnyCollection any](arr []AnyCollection, predicate func(AnyCollection) bool) []AnyCollection {
	var res []AnyCollection
	for _, v := range arr {
		if !predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
