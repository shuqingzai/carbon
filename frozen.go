package carbon

import "sync"

// FrozenNow defines a FrozenNow struct.
// 定义 FrozenNow 结构体
type FrozenNow struct {
	isFrozen bool
	testNow  *Carbon
	rw       *sync.RWMutex
}

var frozenNow = &FrozenNow{
	rw: new(sync.RWMutex),
}

// SetTestNow sets a test Carbon instance for now, remember to clear after used.
// 设置当前测试时间，使用完别忘清除
func SetTestNow(carbon *Carbon) {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = true
	frozenNow.testNow = carbon
}

// CleanTestNow clears a test Carbon instance for now.
// 清除当前测试时间
func CleanTestNow() {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = false
}

// IsTestNow reports whether is testing time.
// 是否是测试时间
func IsTestNow() bool {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	return frozenNow.isFrozen
}
