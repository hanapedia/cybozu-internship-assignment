package worker

import (
	"sync"

	"github.com/hanapedia/cybozu-internship-assignment/pkg/hash"
	"github.com/hanapedia/cybozu-internship-assignment/pkg/models"
)

// worker receives each line and returns processed line.
func worker(jobs <-chan models.Job, results chan<- models.Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- models.Result{
			Index: j.Index,
			Hash:  hash.Hash(j),
		}
	}
}

// SpawnWorkers spawns workers with predefined channels and wait group
// numWorkers is determined by the product of the number of physical cpu and multiplier.
func SpawnWorkers(jobs <-chan models.Job, results chan<- models.Result, wg *sync.WaitGroup, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, wg)
	}
}
