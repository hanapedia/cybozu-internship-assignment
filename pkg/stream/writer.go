package stream

import (
	"fmt"
	"sync"

	"github.com/hanapedia/cybozu-internship-assignment/pkg/models"
)

// WriteStream outputs the processed lines in the original order.
func WriteStream(results <-chan models.Result, wg *sync.WaitGroup) {
	defer wg.Done()
	// buffer is a map to buffer the processed line so that the original order is preserved.
	buffer := make(map[int]string)
	nextIndex := 0

	for res := range results {
		buffer[res.Index] = res.Hash

		// check and wait for the next line in order is ready to be printed.
		for hash, ok := buffer[nextIndex]; ok; hash, ok = buffer[nextIndex] {
			fmt.Println(hash)
			// delete the map entry after printed
			delete(buffer, nextIndex)
			nextIndex++
		}
	}
}
