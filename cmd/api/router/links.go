package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"urlstore/cmd/api/resource/database"
	"urlstore/cmd/api/resource/model"
)

type Api struct {
	D *database.Postgres
}

func genId(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func (a *Api) Create(w http.ResponseWriter, r *http.Request) {
	l := new(model.Link)
	err := json.NewDecoder(r.Body).Decode(&l)

	var j []byte

	if err != nil {
		m := model.Error{
			Code:    strconv.Itoa(http.StatusUnprocessableEntity),
			Message: err.Error(),
			Header:  http.StatusUnprocessableEntity,
		}

		m.Print(w)

		return
	}

	m, err := regexp.MatchString("https?:\\/\\/?\\w+\\.\\w+", l.Url)

	if m == false {
		m := model.Error{
			Code:    strconv.Itoa(http.StatusUnprocessableEntity),
			Message: "invalid url",
			Header:  http.StatusUnprocessableEntity,
		}

		m.Print(w)
		return
	}

	l.Id = genId(5)
	l.CreatedAt = time.Now()

	err = a.D.Insert(l)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			m := model.Error{
				Code:    pgErr.Code,
				Message: pgErr.Message,
				Header:  http.StatusConflict,
			}

			m.Print(w)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	j = []byte(`{"status":"OK", "id":"` + l.Id + `"}`)
	w.Write(j)
}

// /{link}

func (a *Api) Fetch(w http.ResponseWriter, r *http.Request) (model.Link, error) {
	l := chi.URLParam(r, "link")

	link, err := a.D.Find(l)

	if link.Id == "" {
		m := model.Error{
			Code:    strconv.Itoa(http.StatusNotFound),
			Message: "link not found",
			Header:  http.StatusNotFound,
		}

		m.Print(w)
	}

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			m := model.Error{
				Code:    pgErr.Code,
				Message: pgErr.Message,
				Header:  http.StatusInternalServerError,
			}

			m.Print(w)
		}
		return model.Link{}, err
	}

	j, err := json.Marshal(link)

	if err != nil {
		m := model.Error{
			Code:    strconv.Itoa(http.StatusUnprocessableEntity),
			Message: err.Error(),
			Header:  http.StatusUnprocessableEntity,
		}

		m.Print(w)

		return model.Link{}, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	return link, nil
}

func (a *Api) Redirect(w http.ResponseWriter, r *http.Request) {
	l := chi.URLParam(r, "link")

	link, _ := a.D.Find(l)

	fmt.Println(link.Url)

	http.Redirect(w, r, link.Url, http.StatusFound)
}
