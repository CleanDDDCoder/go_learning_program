package pprof

import (
	"os"
	"runtime/pprof"
)

// ProfileCPUFixed starts CPU profiling for the duration the returned stop function is not called.
func ProfileCPUFixed(filename string) (stop func() error) {
	f, err := os.Create(filename)
	if err != nil {
		return func() error { return err }
	}
	pprof.StartCPUProfile(f)
	return func() error {
		pprof.StopCPUProfile()
		return f.Close()
	}
}

// ProfileHeapFixed captures and writes a heap profile to the specified file.
func ProfileHeapFixed(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return pprof.WriteHeapProfile(f)
}