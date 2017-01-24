package baseutil

const ALPHABET_SHIFT  = 87
const NUMERICAL_SHIFT = 48

func charIndex(b byte) byte {
	if b > 9 {
		return b + ALPHABET_SHIFT
	}

	return b + NUMERICAL_SHIFT
}

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
