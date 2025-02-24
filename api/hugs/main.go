package handler

import (
	"net/http"

	"github.com/lucialv/anime-api-cdn/endpoints"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	endpoints.HugsHandler(w, r)
}
