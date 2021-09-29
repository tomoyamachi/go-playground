package main

import (
	"log"
	"time"
)

func main() {
	timeoutOrPanic()
}

func timeoutOrPanic() {
	pChan := make(chan struct{})
	done := make(chan struct{})
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		defer func() {
			// panicが起こればpChanに連絡
			p := recover()
			log.Printf("recover %v\n", p)
			if p != nil {
				log.Printf("panic description %v\n", p)
				pChan <- struct{}{}
			}
			// waitAndPanicが完了し、panicが起こってなければdoneに信号を送る
			close(done)
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
	close(pChan)
	ticker.Stop()
	log.Println("finish")
}

func waitAndPanic() {
	time.Sleep(time.Second * 3)
	panic("test")
}
