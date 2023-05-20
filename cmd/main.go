package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/hanapedia/cybozu-internship-assignment/pkg/models"
	"github.com/hanapedia/cybozu-internship-assignment/pkg/stream"
	"github.com/hanapedia/cybozu-internship-assignment/pkg/worker"
)

func main() {
	// Define CLI flag
	filePath := flag.String("file", "", "File path to process")
	multiplier := flag.Int("multiplier", 1, "Multiplier for the number of goroutines, relative to the number of CPU cores")
	flag.Parse()

	// Check if file path is provided
	if *filePath == "" {
		fmt.Println("Please provide a file path with the -file flag")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Create buffered reader
	scanner := bufio.NewScanner(file)

	// Create channels and worker pool
	jobs := make(chan models.Job)
	results := make(chan models.Result)
	var workerWaitGroup sync.WaitGroup
	// Calculate number of workers based on CPU cores and multiplier
	numWorkers := runtime.NumCPU() * *multiplier
	// Spawn the workers
	worker.SpawnWorkers(jobs, results, &workerWaitGroup, numWorkers)

	// Print results in order
	// make sure that the program does not exit before printing is complete
	var printWaitGroup sync.WaitGroup
	printWaitGroup.Add(1)
	go stream.WriteStream(results, &printWaitGroup)

	// Read and dispatch jobs
	stream.ReadStream(scanner, jobs)

	// Check scanner error
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wait for all workers to finish then close results channel
	go func() {
		workerWaitGroup.Wait()
		close(results)
	}()

	// Wait for printing to complete
	printWaitGroup.Wait()
}
