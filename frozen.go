package carbon

import "sync"

// TestNow defines a TestNow struct.
// 定义 TestNow 结构体
type TestNow struct {
	isFrozen  bool
	frozenNow *Carbon
	rw        *sync.RWMutex
}

var testNow = &TestNow{
	rw: new(sync.RWMutex),
}

// SetTestNow sets a test Carbon instance for now, remember to clear after used.
// 设置当前测试时间，使用完别忘清除
func SetTestNow(carbon *Carbon) {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	testNow.isFrozen = true
	testNow.frozenNow = carbon
}

// CleanTestNow clears a test Carbon instance for now.
// 清除当前测试时间
func CleanTestNow() {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	testNow.isFrozen = false
}

// IsTestNow reports whether is testing time.
// 是否是测试时间
func IsTestNow() bool {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	return testNow.isFrozen
}
