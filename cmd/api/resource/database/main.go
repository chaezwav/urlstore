package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"urlstore/cmd/api/resource/model"
	"urlstore/config"
)

type Postgres struct {
	d *pgxpool.Pool
	c context.Context
}

var p *Postgres

func LoadDatabase(ctx context.Context) *Postgres {
	godotenv.Load()
	c := config.LoadDatabaseConfig()

	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User.Name, c.User.Password, c.Name)

	d, err := pgxpool.New(ctx, s)

	if err != nil {
		fmt.Errorf("couldn't connect to the database: %w", err)
	}

	p = &Postgres{d, ctx}
	return p
}

func (p *Postgres) Ping() error {
	err := p.d.Ping(p.c)

	if err != nil {
		return fmt.Errorf("error: database: %s", err.Error())
	} else {
		fmt.Println("database: pinged!")
	}

	return nil
}

func (p *Postgres) Insert(data *model.Link) error {
	r, err := p.d.Exec(p.c, `INSERT INTO links VALUES ($1, $2, $3, $4, $5)`, data.Id, data.Alias, data.Url, data.Flags, data.CreatedAt)

	if err != nil {
		return fmt.Errorf("error: database: %w", err)
	} else {
		fmt.Printf("database: inserted link %v %s", data.Id, r)
	}

	return nil

}

func (p *Postgres) Find(key string) (model.Link, error) {

	r, err := p.d.Query(p.c, `SELECT * FROM links WHERE (id = $1 OR alias = $1)`, key)

	if err != nil {
		fmt.Errorf("error: database: %w", err)
	}

	link, err := pgx.CollectOneRow(r, pgx.RowToStructByPos[model.Link])

	if err != nil {
		fmt.Errorf("error: database: %w", err)
	} else {
		fmt.Printf("database: found link %v", link.Id)
	}

	return link, nil
}
