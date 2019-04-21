package channels

import (
	"time"
	"sync"
)

//ChannelProtection on a data string
type ChannelProtection struct {
	data string
	writeData chan string
	readRequest chan struct{}
	readData chan string
	shutdown chan struct{}
	wg *sync.WaitGroup
} 
//NewChannelProtection ...
func NewChannelProtection() *ChannelProtection {
	return &ChannelProtection{
		writeData: make(chan string, 10),
		readRequest: make(chan struct{}, 10),
		readData: make(chan string, 10),
		shutdown: make(chan struct{}, 10),
		wg: &sync.WaitGroup{},
	}
}
//Read the current data
// this function runs in goroutine that calls it
func (m ChannelProtection) Read() string {
	m.readRequest <- struct{}{}
	currentData := <- m.readData
	return currentData
}
//Write something to the data
// this function runs in the goroutine that calls it
func (m ChannelProtection) Write(s string) {
	m.writeData <- s
}
//write is done in a different goroutine and the data in the input 
// no longer "belongs" to the sender.  This means that if it is a pointer 
// type like a map or a slice the sender must send a copy
func (m *ChannelProtection) write(s string) {
	m.data = s
	time.Sleep(10 * time.Millisecond) // Simulates "work" done to save the data
}
//read is done in a different gorioutine and the data returned 
// is "given" to the caller of Read through a channel and must be a 
// copy of the data at a given time and cannot reference internal data
func (m ChannelProtection) read() {
	time.Sleep(10 * time.Millisecond) // Simulates "work" done to retrieve the data
	m.readData <- "test"
}
func (m ChannelProtection) background() {
	mainloop:
	for {
		select {
		case <- m.shutdown:
			for {
				select {
				case newString := <- m.writeData:
					m.write(newString)
				case <- m.readRequest:
					go m.read()
				default:
					break mainloop
				}
			}
		case newString := <- m.writeData:
			m.write(newString)
		case <- m.readRequest:
			go m.read()
		}
	}
	m.wg.Done()
}

//Start the background processing
func (m ChannelProtection) Start() {
	m.wg.Add(1)
	go m.background()
}
//Close down the background resources
func (m ChannelProtection) Close() {
	m.shutdown <- struct{}{}
	m.wg.Wait()
}