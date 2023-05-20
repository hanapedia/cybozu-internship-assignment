package profiling

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func RunMemoryProfiling() {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal(err)
	}
}
