package concurrency

import (
	"fmt"
	"time"
)

// A code for synchrous communicating
// "count" is  a producer, while msg is a consumer.
// count will be blocked until find a consumer.
func Synccom() {
	c := make(chan string)
	go count("sleep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
