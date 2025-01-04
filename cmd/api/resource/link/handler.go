package link

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"slices"
	"strconv"
)

type Link struct {
	ID    int      `json:"ID"`
	Url   string   `json:"url"`
	Flags []string `json:"flags"`
}

var links = []Link{
	{ID: 1, Url: "https://koehn.lol", Flags: []string{"ephemeral"}},
	{ID: 2, Url: "https://social.lol", Flags: []string{"ephemeral"}},
}

func findLink(i int) (Link, error) {
	idx := slices.IndexFunc(links, func(l Link) bool { return l.ID == i })

	if idx < 0 {
		return Link{}, errors.New("not found")
	}

	return links[idx], nil
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(r, "id"))

	link, err := findLink(id)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	http.Redirect(w, r, link.Url, http.StatusSeeOther)
	fmt.Println("redirect: ", link.Url)
}

func Fetch(w http.ResponseWriter, r *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(r, "id"))
	link, err := findLink(id)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(link)
	fmt.Println("fetch: ", link.Url)
}
