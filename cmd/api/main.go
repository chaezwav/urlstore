package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"urlstore/cmd/api/resource/link"
	"urlstore/cmd/api/router"
	"urlstore/config"
)

func main() {
	godotenv.Load()
	c := config.LoadConfig()

	if c.Debug == true {
		log.Println("[ DEBUG ENVIRONMENT ]")
	}

	_ = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DB.Host, c.DB.Port, c.DB.User.Username, c.DB.User.Password, c.DB.DBName)
	//d, err := pgxpool.New(context.Background(), dbs)
	//
	//p := &link.Postgres{Db: d, Ctx: context.Background()}

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", c.Server.Port),
		//Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Printf("init: starting server on %s", s.Addr)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("error: init: ", err.Error())
	}
}
