package sieve

// find and return prime numbers as list between [2,limit]
func Sieve(limit int) []int{
	var primeList []int
	var primeCheck = make([]int, limit + 1)
	if limit == 2{
		return append(primeList, 2)
	}
	i := 2
	for i < len(primeCheck) {
		currentVal := i + i
		for currentVal <= limit{
			primeCheck[currentVal] = 1
			currentVal += i
		}

		// increment i
		j := i+1
		for j <= limit &&  primeCheck[j] == 1{
			j++
		}
		i = j
	}

	// add not signed numbers
	for i := 2; i < len(primeCheck) ; i++{
		if primeCheck[i] != 1 {
			primeList = append(primeList, i)
		}
	}

	return primeList
}




