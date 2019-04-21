package channels

import (
	"sync"
	"testing"
)

func benchMutex(workers int, workload int, b *testing.B) {
	m := NewMutextProtection()
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	writer := func() {
		for i := 0; i < workload; i++ {
			m.Write("test")
		}
		wg.Done()
	}
	reader := func() {
		for i := 0; i < workload; i++ {
			_ = m.Read()
		}
		wg.Done()
	}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go writer()
		wg.Add(1)
		go reader()
	}
	wg.Wait()
}
func BenchmarkMutexWorks1by1000(b *testing.B) {
	benchMutex(1, 1000, b)
}
func BenchmarkMutexWorks10by100(b *testing.B) {
	benchMutex(10, 100, b)
}
func BenchmarkMutexWorks100by10(b *testing.B) {
	benchMutex(100, 10, b)
}
func BenchmarkMutexWorks1000by1(b *testing.B) {
	benchMutex(1000, 1, b)
}