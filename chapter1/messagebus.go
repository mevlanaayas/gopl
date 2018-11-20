package main

import (
	"fmt"
	"github.com/vardius/message-bus"
	"sync"
)

/*

in this file We will see how to import third party libraries
simply run next statement
	go get github.com/vardius/message-bus
now message-bus will be placed under GOPATH location and will be importable.

*/

func main() {
	bus := messagebus.New(10)

	var wg sync.WaitGroup
	wg.Add(2)

	bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("üst ", v)
	})

	bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("alt ", v)
	})

	bus.Publish("topic", true)
	wg.Wait()
}
