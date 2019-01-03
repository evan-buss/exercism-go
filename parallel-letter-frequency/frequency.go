package letter

import (
	"sync"
)

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

// ConcurrentFrequency creates a thread for each string in s
// Returns a combined map of all the frequencies
func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	var mux = &sync.Mutex{}
	var wg sync.WaitGroup
	defer wg.Wait() //Wait at the end

	for _, str := range s {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			for _, r := range text {
				mux.Lock()
				m[r]++
				mux.Unlock()
			}
		}(str)
	}
	return m
}
