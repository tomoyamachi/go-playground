package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	go timeoutOrPanic()
	time.Sleep(time.Minute)
}

func timeoutOrPanic() {
	pChan := make(chan struct{})
	done := make(chan struct{})
	ticker := time.NewTicker(3 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			// panicが起こればpChanに連絡
			p := recover()
			log.Printf("recover %v\n", p)
			if p != nil {
				log.Printf("panic description: %v\n", p)
				pChan <- struct{}{}
			}
			// waitAndPanicが完了し、panicが起こってなければdoneに信号を送る
			close(done)
			log.Println("close doneChan")

		}()
		waitAndPanic()
	}()
	log.Println("start select")
	select {
	case <-pChan:
		log.Println("panic occured")
	case <-done:
		log.Println("done occured")
	case <-ticker.C:
		log.Println("timeout occured")
	}
	ticker.Stop()
	wg.Wait()
	close(pChan)
	log.Println("finish")
}

func waitAndPanic() {
	time.Sleep(time.Second * 5)
	panic("test")
}
