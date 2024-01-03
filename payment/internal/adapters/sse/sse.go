package sse

import (
	"encoding/json"
	"fmt"
	"github.com/pwera/grpc-micros-payment/internal/application/core/domain"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Adapter struct {
	m        *sync.Mutex
	counter  int64
	requests map[int64]chan domain.Payment
	mux      http.ServeMux
}

func NewAdapter() *Adapter {
	adapter := &Adapter{
		m:        new(sync.Mutex),
		requests: map[int64]chan domain.Payment{},
		mux:      http.ServeMux{},
	}
	adapter.mux.HandleFunc("/sse", adapter.SSE)
	return adapter
}

func (s *Adapter) SSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	id := atomic.AddInt64(&s.counter, 1)
	events := make(chan domain.Payment)
	s.requests[id] = events
	defer func() {
		s.m.Lock()
		defer s.m.Unlock()
		delete(s.requests, id)
		close(events)
	}()

	timer := time.NewTimer(0)
loop:
	for {
		select {
		case <-timer.C:
			if _, err := fmt.Fprintf(w, "event: message\ndata: ping\n\n"); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			timer.Reset(time.Second * 5)
		case e := <-events:
			fmt.Println("Sending reload event...")
			js, _ := json.Marshal(&e)
			if _, err := fmt.Fprintf(w, "event: \ndata: %s\n\n", string(js)); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case <-r.Context().Done():
			break loop
		}
		w.(http.Flusher).Flush()
	}
}

func (s *Adapter) Send(p domain.Payment) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, f := range s.requests {
		f := f
		go func(f chan domain.Payment) {
			f <- p
		}(f)
	}
}

func (s *Adapter) Run() {
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 9999), &s.mux); err != nil {
			fmt.Printf("Error starting proxy: %v\n", err)
		}
	}()
}
