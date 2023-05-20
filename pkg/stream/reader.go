package stream

import (
	"bufio"

	"github.com/hanapedia/cybozu-internship-assignment/pkg/models"
)

// ReadStream reads the file lines one by one and pass it to worker pool
func ReadStream(scanner *bufio.Scanner, jobs chan<- models.Job) {
	index := 0
	for scanner.Scan() {
		jobs <- models.Job{
			Index: index,
			Line:  scanner.Text(),
		}
		index++
	}
	close(jobs)
}
