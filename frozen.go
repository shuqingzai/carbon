package carbon

import "sync"

// FrozenNow defines a FrozenNow struct.
type FrozenNow struct {
	isFrozen bool
	testNow  *Carbon
	rw       *sync.RWMutex
}

var frozenNow = &FrozenNow{
	rw: new(sync.RWMutex),
}

// SetTestNow sets a test Carbon instance for now.
func SetTestNow(c *Carbon) {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = true
	frozenNow.testNow = c
}

// ClearTestNow clears the test Carbon instance for now.
func ClearTestNow() {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = false
	frozenNow.testNow = nil
}

// IsTestNow reports whether is testing time.
func IsTestNow() bool {
	frozenNow.rw.RLock()
	defer frozenNow.rw.RUnlock()

	return frozenNow.isFrozen
}
