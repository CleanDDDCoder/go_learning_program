//go:build ignore

package maps

// InventoryTotal sums item counts in an inventory map.
func InventoryTotal(inventory map[string]int) int {
	total := 0
	for _, count := range inventory {
		total += count
	}
	return total
}

// HasItem reports whether inventory contains item.
func HasItem(inventory map[string]int, item string) bool {
	_, ok := inventory[item]
	return ok
}
