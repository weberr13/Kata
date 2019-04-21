package channels

import (
	"sync"
	"testing"
)


func benchChannel(workers int, workload int,  b *testing.B) int {
	m := NewChannelProtection()
	m.Start()
	wg := &sync.WaitGroup{}
	b.ResetTimer()

	writer := func () {
		for i := 0 ; i < workload ; i++ {
			m.Write("test")
		}
		wg.Done()
	}
	reader := func () {
		for i := 0 ; i < workload ; i++ {
			_ = m.Read()
		}
		wg.Done()
	}
	for i := 0 ; i < workers ; i++ {
		wg.Add(1)
		go writer()
		wg.Add(1)
		go reader()
	}
	wg.Wait()
	m.Close()
	return 0
}
func BenchmarkChannelsWorkers1by1000(b *testing.B) {
	_ = benchChannel(1, 1000, b)
}
func BenchmarkChannelsWorkers10by100(b *testing.B) {
	_ = benchChannel(10, 100, b)
}
func BenchmarkChannelsWorkers100by10(b *testing.B) {
	_ = benchChannel(100, 10, b)
}
func BenchmarkChannelsWorkers1000by1(b *testing.B) {
	_ = benchChannel(1000, 1, b)
}