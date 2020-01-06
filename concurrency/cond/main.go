package main

import (
	"log"
	"sync"
	"time"
)

type EventSendMessage struct {
	*sync.Cond
}

func main() {

	e := EventSendMessage{
		sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	subscribe(e.Cond, func() {
		log.Print("message sended")
	})

	e.Cond.Broadcast()
	e.Cond.Broadcast()

	time.Sleep(1 * time.Second)

	go func(e *EventSendMessage) {
		//EventSendMessage{sync.NewCond(&sync.Mutex{})}.Broadcast()
		e.Broadcast()
	}(&e)

	time.Sleep(1 * time.Second)
}
