package maps

import "testing"

func TestInventoryTotal(t *testing.T) {
	inventory := map[string]int{"apples": 3, "oranges": 5}
	if got := InventoryTotal(inventory); got != 8 {
		t.Fatalf("InventoryTotal() = %d, want 8", got)
	}
}

func TestHasItem(t *testing.T) {
	inventory := map[string]int{"apples": 0}
	if !HasItem(inventory, "apples") {
		t.Fatal("HasItem() = false for existing zero-count item")
	}
	if HasItem(inventory, "bananas") {
		t.Fatal("HasItem() = true for missing item")
	}
}
