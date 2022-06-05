package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutex sync.Mutex

var ticket = 10

func main() {
	wg.Add(4)
	go saleTicket("售票窗口1")
	go saleTicket("售票窗口2")
	go saleTicket("售票窗口3")
	go saleTicket("售票窗口4")
	wg.Wait()
}

func saleTicket(name string) {
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Second)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售完，没有票了：", ticket)
			break
		}
		mutex.Unlock()
	}
}
