package models

// job represents each line from file that is passed to the worker to be processed.
type Job struct {
	// index is the index of the line in the file.
	Index int

	// line is the content of the line in string.
	Line  string
}

// result represents each line after it was processed by the worker
type Result struct {
	// index is the index of the processed line in the original file.
	Index int

	// hash is the content of the lien after it was processed
	Hash  string
}

