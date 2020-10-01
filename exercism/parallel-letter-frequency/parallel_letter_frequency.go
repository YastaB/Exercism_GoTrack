package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text array concurrently and returns this
// data as a FreqMap.
func ConcurrentFrequency(s []string) FreqMap {
	rtnFreqMap := FreqMap{}
	freqMapChan := make(chan FreqMap, 10)

	for _, text := range s {
		go func(text string) {
			freqMapChan <- Frequency(text)
		}(text)
	}

	// read from the buffered channel and apply to the return struct
	for range s {
		resFreqMap := <-freqMapChan
		for key, val := range resFreqMap {
			rtnFreqMap[key] += val
		}
	}

	return rtnFreqMap
}
