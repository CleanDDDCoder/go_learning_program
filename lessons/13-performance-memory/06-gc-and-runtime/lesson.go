package gcruntime

import (
	"runtime"
	"runtime/debug"
)

// GCSettings represents configurable GC settings.
type GCSettings struct {
	// GOGC controls the heap growth percentage.
	GOGC int
	// GOMEMLIMIT sets the memory limit in bytes.
	GOMEMLIMIT int64
}

// GetGOGC returns the current GOGC value.
//
// TODO: Implement using runtime.GOGC.
func GetGOGC() int {
	return debug.SetGCPercent(-1) // -1 returns current value without setting
}

// SetGOGC sets the GOGC value.
//
// TODO: Implement using debug.SetGOGC.
func SetGOGC(value int) error {
	debug.SetGCPercent(value)
	return nil
}

// GetMemStats returns current memory statistics.
//
// TODO: Implement using runtime.ReadMemStats.
func GetMemStats() *runtime.MemStats {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	return stats
}

// GC triggers a garbage collection run.
//
// TODO: Implement using runtime.GC.
func GC() {
	runtime.GC()
}

// NumCPU returns the number of logical CPUs.
//
// TODO: Implement using runtime.NumCPU.
func NumCPU() int {
	return runtime.NumCPU()
}

// NumGoroutine returns the number of goroutines.
//
// TODO: Implement using runtime.NumGoroutine.
func NumGoroutine() int {
	return runtime.NumGoroutine()
}
