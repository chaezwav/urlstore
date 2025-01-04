package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"urlstore/appconfig"
	"urlstore/cmd/api/router"
)

func main() {
	cfg, err := appconfig.LoadFromPath(context.Background(), "cmd/api/config/local/config.pkl")

	if err != nil {
		panic(err)
	}

	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Timeout.ReadDuration) * time.Second,
		WriteTimeout: time.Duration(cfg.Timeout.WriteDuration) * time.Second,
		IdleTimeout:  time.Duration(cfg.Timeout.IdleDuration) * time.Second,
	}

	if cfg.Debug == true {
		log.Println("[ DEBUG ENVIRONMENT ]")
	}

	log.Printf("init: starting server on %s", s.Addr)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("error: init: ", err)
	}
}
