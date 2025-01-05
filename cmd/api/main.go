package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"urlstore/cmd/api/resource/database"
	"urlstore/cmd/api/router"
	"urlstore/config"
)

func main() {
	godotenv.Load()
	c := config.LoadConfig()

	if c.Debug == true {
		log.Println("[ DEBUG ENVIRONMENT ]")
	}

	d := database.LoadDatabase(context.Background())

	r := router.New(d)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Printf("init: starting server on %s", s.Addr)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("error: init: ", err.Error())
	}
}
