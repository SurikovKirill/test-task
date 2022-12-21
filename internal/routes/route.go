package routes

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"test-task/internal/exchange"
)

// @Summary Get Exchange rate for currency
// @Description Get Exchange
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /exchange/{currency} [get]
func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("X-API-KEY") != "123321" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Println(r.URL.Path)
	a, err := exchange.GetExchange(curVal(strings.Split(r.URL.Path, "/")))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func curVal(tmp []string) string {
	if tmp[2] == "" {
		var cur = [5]string{"BGN", "BRL", "HUF", "HKD", "DKK"}
		return cur[rand.Intn(5)]
	} else {
		return tmp[len(tmp)-1]
	}
}
