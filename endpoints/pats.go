package endpoints

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/lucialv/anime-api-cdn/common"
)

func PatsHandler(w http.ResponseWriter, r *http.Request) {
	totalGifs := 30
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(totalGifs) + 1

	gifURL := fmt.Sprintf("https://cdn.lucia-dev.com/pat%02d.gif", randomNumber)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(common.Response{URL: gifURL})
}
