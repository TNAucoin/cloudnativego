package main

import (
	"fmt"
	"github.com/tnaucoin/cloudnativego/cmd/api/router"
	"github.com/tnaucoin/cloudnativego/config"
	"io"
	"log"
	"net/http"
)

func main() {
	c := config.New()
	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}
	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("server startup failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}
