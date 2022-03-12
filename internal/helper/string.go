package helper

// PadLeft adds padding string to the left of a string until the padded string reach the specified max length.
func PadLeft(str, pad string, maxLength int) string {
	if pad != "" {
		for len(str)+len(pad) <= maxLength {
			str = pad + str
		}
	}
	return str
}

// PadRight adds padding string to the right of a string until the padded string reach the specified max length.
func PadRight(str, pad string, maxLength int) string {
	if pad != "" {
		for len(str)+len(pad) <= maxLength {
			str = str + pad
		}
	}
	return str
}
