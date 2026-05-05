//go:build ignore

package methods

// Rectangle stores width and height values.
type Rectangle struct {
	Width  int
	Height int
}

// Area returns the rectangle area.
func (rect Rectangle) Area() int {
	return rect.Width * rect.Height
}

// Perimeter returns the distance around the rectangle.
func (rect Rectangle) Perimeter() int {
	return 2 * (rect.Width + rect.Height)
}
