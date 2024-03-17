package factor

func Factor_Number(primes []int, value int) (output []int) {

	for i := 0; i < len(primes); {
		if value%primes[i] == 0 {
			output = append(output, primes[i])
			value = value / primes[i]
		} else {
			i++
		}
	}
	if value > 1 {
		output = append(output, value)
	}

	return output
}
