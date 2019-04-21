package channels

import (
	"time"
	"sync"
)

//MutexProtection on a data string
type MutexProtection struct {
	data string
	lock *sync.Mutex
} 
//NewMutextProtection ...
func NewMutextProtection() *MutexProtection {
	return &MutexProtection{lock: &sync.Mutex{}}
}
//Read the current data
func (m MutexProtection) Read() string {
	m.lock.Lock()
	time.Sleep(10*time.Millisecond) // simulates "work" done to retrieve the data
	m.lock.Unlock()

	return "test"
}
//Write something to the data
func (m *MutexProtection) Write(s string) {
	m.lock.Lock()
	m.data = s
	time.Sleep(10*time.Millisecond) // simultes "work" done to save the data
	m.lock.Unlock()
}