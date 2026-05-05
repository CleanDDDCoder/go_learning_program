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

// GetGOGCFixed returns the current GOGC value.
func GetGOGCFixed() int {
	return debug.SetGCPercent(-1) // -1 returns current value without setting
}

// SetGOGCFixed sets the GOGC value.
func SetGOGCFixed(value int) error {
	debug.SetGCPercent(value)
	return nil
}

// GetMemStatsFixed returns current memory statistics.
func GetMemStatsFixed() *runtime.MemStats {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	return stats
}

// GCFixed triggers a garbage collection run.
func GCFixed() {
	runtime.GC()
}

// NumCPUFixed returns the number of logical CPUs.
func NumCPUFixed() int {
	return runtime.NumCPU()
}

// NumGoroutineFixed returns the number of goroutines.
func NumGoroutineFixed() int {
	return runtime.NumGoroutine()
}