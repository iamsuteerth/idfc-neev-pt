package typeconversions

/*
Some relevant theory

Interfaces in Go can introduce ambiguity about the underlying type.
A type assertion allows us to extract the interface value's underlying
concrete value using this syntax: interfaceVariable.(concreteType)

var input interface{} = 12
number := input.(int)

str, ok := input.(string) // no panic if input is not a string

A type switch can perform several type assertions in series.
It has the same syntax as a type assertion (interfaceVariable.(concreteType)),
but the concreteType is replaced with the keyword type.

var i interface{} = 12 // try: 12.3, true, int64(12), []int{}, map[string]int{}

switch v := i.(type) {
case int:
    fmt.Printf("the integer %d\n", v)
case string:
    fmt.Printf("the string %q\n", v)
default:
    fmt.Printf("type, %T, not handled explicitly: %#v\n", v, v)
}
*/
import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNumber, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	num, _ := strconv.Atoi(fancyNumber.n)
	return num
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fancyNumber, ok := fnb.(FancyNumber)
	if !ok {
		return "This is a fancy box containing the number 0.0"
	}
	num, _ := strconv.Atoi(fancyNumber.n)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))

}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
