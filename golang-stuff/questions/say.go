package questions

var parts = []struct {
	value int64
	name  string
}{
	{1000000000, "billion"},
	{1000000, "million"},
	{1000, "thousand"},
	{100, "hundred"},
	{90, "ninety"},
	{80, "eighty"},
	{70, "seventy"},
	{60, "sixty"},
	{50, "fifty"},
	{40, "forty"},
	{30, "thirty"},
	{20, "twenty"},
	{19, "nineteen"},
	{18, "eighteen"},
	{17, "seventeen"},
	{16, "sixteen"},
	{15, "fifteen"},
	{14, "fourteen"},
	{13, "thirteen"},
	{12, "twelve"},
	{11, "eleven"},
	{10, "ten"},
	{9, "nine"},
	{8, "eight"},
	{7, "seven"},
	{6, "six"},
	{5, "five"},
	{4, "four"},
	{3, "three"},
	{2, "two"},
	{1, "one"},
}

// Say converts the given integer to an english phrase
func Say(n int64) (string, bool) {
	if n < 0 || n >= 1e12 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	return sayPart(n, ""), true
}

func sayPart(n int64, prefix string) string {
	// Going from top to bottom
	for _, p := range parts {
		// value is 1000 and name is "thousand"
		// n = 1056 but p.value = 1000000 will make order as 0 and rem as n
		order, rem := n/p.value, n%p.value
		if order == 0 {
			// The order of magnitude is not matching
			continue
		}
		if p.value < 100 {
			// Role of prefix is to add the space or the hyphen
			// If we are at 90, then p.name is ninety and the prefix is going to be hyphen for the remaining part, like 8 which will be inside rem
			return prefix + p.name + sayPart(rem, "-")
		}
		// Add the order like thousand which is 1-9 giving us results like nine thousand and then for the remaining section
		// we have the sayPart(rem, " ") which basically makes us add a space in the front
		return sayPart(order, prefix) + " " + p.name + sayPart(rem, " ")
	}
	/*
	   Dry RUN for 19876
	   19876 / 1000 gives order = 19 and rem = 876
	   	sayPart(19, "") + " " + thousand + sayPart(876, " ")
	       	"19 thousand " + sayPart(876, " ")
	   876 / 100 gives order = 8 and rem = 76
	       sayPart(8, " ") + hundred + sayPart(76, " ")
	       	"19 thousand 8 hundred " + sayPart(76, " ")
	*/
	return ""
}
