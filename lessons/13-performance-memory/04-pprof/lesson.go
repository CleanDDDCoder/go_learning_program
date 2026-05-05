package pprof

import (
	"os"
	"runtime/pprof"
)

// ProfileCPU starts CPU profiling and returns a function to stop it.
//
// TODO: Implement CPU profiling using pprof.StartCPUProfile.
// TODO: The returned function should call pprof.StopCPUProfile.
func ProfileCPU(filename string) (stop func() error) {
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

// ProfileHeap captures a heap profile and writes to filename.
//
// TODO: Implement heap profiling using pprof.WriteHeapProfile.
func ProfileHeap(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return pprof.WriteHeapProfile(f)
}
