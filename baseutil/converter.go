package baseutil

const (
	ALP_SHIFT byte = 87
	NUM_SHIFT byte = 48
)

type (
	ANSI byte
	Numb10 int64
	BasedNumb struct {
		Base byte
		Numb string
	}
)

func (numb10 Numb10) BaseOf(base byte) BasedNumb {
	var result []byte
	var number int64 = int64(numb10)
	for number > 0 {
		var rank byte = byte(number % int64(base))
		result = append(result, ANSI(rank).IndexOf())

		number = number / int64(base)
		if number < int64(base) {
			result = append(result, ANSI(byte(number)).IndexOf())
			break
		}
	}

	for i, j := 0, len(result)-1; i < len(result)/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return BasedNumb{
		Base: base,
		Numb: string(result),
	}
}

func (b ANSI) IndexOf() byte {
	if b > 9 {
		return byte(b) + ALP_SHIFT
	}

	return byte(b) + NUM_SHIFT
}