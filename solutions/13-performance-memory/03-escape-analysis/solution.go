package escape

// StackAllocFixed shows the correct approach - return value directly to avoid heap allocation.
func StackAllocFixed(n int) int {
	return n + 1
}

// NoEscapeFixed demonstrates pre-allocating slice capacity to avoid escapes.
func NoEscapeFixed(items []int) int {
	sum := 0
	// Pre-allocate with enough capacity to avoid reallocation during append.
	// However, since we're not appending in this case, this is just for demonstration.
	result := make([]int, 0, len(items))
	for _, v := range items {
		result = append(result, v)
		sum += v
	}
	_ = result // suppress unused warning
	return sum
}

// EscapeViaInterfaceFixed avoids unnecessary interface conversion.
func EscapeViaInterfaceFixed(v string) string {
	// Direct string passing avoids interface allocation
	return v
}

// PointerEscape shows how returning pointers causes heap allocation.
// This cannot be avoided if you need to return a pointer.
func PointerEscape(n int) *int {
	result := n + 1
	return &result
}

// StructWithSlice shows that returning a struct containing a slice
// will cause the slice to escape to the heap.
type DataWithSlice struct {
	Values []int
	Name   string
}

func StructWithSlice(items []int) DataWithSlice {
	return DataWithSlice{
		Values: items,
		Name:   "data",
	}
}