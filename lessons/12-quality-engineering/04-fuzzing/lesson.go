package fuzzing

// Fuzz testing uses randomized inputs to discover bugs.
// Go 1.18+ includes native fuzzing support via go test -fuzz.

func Process(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	return string(data), nil
}

func ValidateEmail(email string) bool {
	if len(email) < 3 {
		return false
	}
	atIdx := -1
	for i, c := range email {
		if c == '@' {
			atIdx = i
			break
		}
	}
	return atIdx > 0 && atIdx < len(email)-1
}