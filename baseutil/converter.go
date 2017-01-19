package baseutil

func Convert(base byte, number10 int64) string {
	var result []byte
	for number10 > 0 {
		var rank byte = byte(number10 % int64(base))
		result = append(result, charIndex(rank))

		number10 = number10 / int64(base)
		if number10 < int64(base) {
			result = append(result, charIndex(byte(number10)))
			break
		}
	}

	for i, j := 0, len(result)-1; i < len(result)/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func charIndex(b byte) byte {
	var shift byte
	if b > 9 {
		shift = 87
	} else {
		shift = 48
	}

	return shift + b
}