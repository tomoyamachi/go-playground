package main

import (
	"log"
	"net/http"
	"time"
)

type response struct {
	code int
	body []byte
}

const timeoutDuration = time.Second * 2
const responseDelay = time.Second * 3

func main() {
	s := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resChan := make(chan *response)
		go child(resChan)
		select {
		case res := <-resChan:
			w.WriteHeader(res.code)
			w.Write(res.body)
		case <-time.Tick(timeoutDuration):
			w.WriteHeader(http.StatusGatewayTimeout)
			w.Write([]byte("強制的にデータを返す"))
		}
	})
	log.Println("curl -v http://localhost:8080/")
	log.Fatal(s.ListenAndServe())
}

func child(ch chan *response) {
	time.Sleep(responseDelay)
	log.Println("子スレッドでは処理を続ける")
	ch <- &response{
		code: http.StatusOK,
		body: []byte("正常にレスポンスを返す"),
	}
}
