package handlers

import (
	"log"
	"io"
	"fmt"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello, world!")
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		
		// write the response back to user
		fmt.Fprintf(rw, "Hello, %s", data)
}
