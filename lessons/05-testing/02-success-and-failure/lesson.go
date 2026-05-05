package success_failure

import "errors"

// ErrDivideByZero is returned when the divisor is zero.
var ErrDivideByZero = errors.New("divide by zero")

// Divide returns integer division or an error for a zero divisor.
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
