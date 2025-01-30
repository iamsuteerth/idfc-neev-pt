package questions

import "errors"

// Define the List and Element types here.
type Element struct {
	Value interface{}
	next  *Element // Changed to *Element for consistency
}

type SList struct {
	head *Element
	tail *Element // Added tail for efficiency in Push
}

func New(elements []int) *SList {
	ll := &SList{}
	for _, v := range elements {
		ll.Push(v)
	}
	return ll
}

func (l *SList) Size() int {
	if l == nil || l.head == nil { // Handle nil list and empty list
		return 0
	}
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *SList) Push(element int) {
	newNode := &Element{Value: element}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode // Initialize tail for the first element
		return
	}
	l.tail.next = newNode // Add to the end of the list using the tail
	l.tail = newNode      // Update tail
}

func (l *SList) Pop() (int, error) {
	if l == nil || l.head == nil {
		return 0, errors.New("empty list")
	}

	val := l.tail.Value.(int) // Type assertion to int.  Handle non-int values if needed.
	if l.head == l.tail {     // Only one element
		l.head = nil
		l.tail = nil
		return val, nil
	}
	current := l.head
	for current.next != l.tail {
		current = current.next
	}
	current.next = nil
	l.tail = current

	return val, nil
}

func (l *SList) Array() []int {
	if l == nil || l.head == nil {
		return []int{} // Return empty slice for nil or empty list
	}
	var arr []int
	current := l.head
	for current != nil {
		arr = append(arr, current.Value.(int)) // Type assertion to int
		current = current.next
	}
	return arr
}

func (l *SList) Reverse() *SList {
	if l == nil || l.head == nil || l.head == l.tail {
		return l // Nothing to reverse
	}

	var prev *Element
	current := l.head
	l.tail = l.head //old head is the new tail
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev //new head
	return l
}
