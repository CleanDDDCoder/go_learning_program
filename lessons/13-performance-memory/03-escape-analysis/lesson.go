package escape

// StackAlloc demonstrates a case where the variable stays on the stack.
// When a value is returned, it must escape to the heap.
//
// TODO: Modify Increment to return the value directly without causing it to escape.
// Consider using a pointer receiver or returning the value itself.
func StackAlloc(n int) *int {
	result := n + 1
	return &result
}

// NoEscape shows variables that do not escape because they are used
// within the function scope only.
//
// TODO: Modify Sum to pre-allocate the slice capacity to avoid escapes.
func NoEscape(items []int) int {
	sum := 0
	for _, v := range items {
		sum += v
	}
	return sum
}

// EscapeViaInterface shows that passing a value to an interface can cause
// it to escape to the heap because interface dispatch requires heap allocation.
//
// TODO: Refactor to avoid the interface conversion, or document why the escape is necessary.
func EscapeViaInterface(v interface{}) string {
	return v.(string)
}

// TODO: Add a function demonstrating pointer escaping via being stored in a slice.
// TODO: Add a function showing how returning a struct with an embedded slice causes escape.
