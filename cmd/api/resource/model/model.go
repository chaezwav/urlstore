package model

import (
	"net/http"
	"strings"
	"time"
)

type Link struct {
	Id        string    `json:"id"`
	Alias     string    `json:"alias,omitempty"`
	Url       string    `json:"url"`
	Flags     []string  `json:"flags,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Header  int    `json:"header"`
}

func (e *Error) Print(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Header)
	w.Write([]byte(`{"code": ` + e.Code + `, "message": "` + strings.ReplaceAll(e.Message, "\"", "'") + `"}`))
}
