package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/jcelliott/turnpike.v2"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type Quote struct {
	Date            int64   `json:"date"`
	Open            float64 `json:"open"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Close           float64 `json:"close"`
	Volume          float64 `json:"volume"`
	QuoteVolume     float64 `json:"quoteVolume"`
	WeightedAverage float64 `json:"weightedAverage"`
}

var currencyPairs = map[string]bool{
	"USDT_BCH":  true, // Bitcoin Cash
	"USDT_BTC":  true, // Bitcoin
	"USDT_DASH": true, // Dash
	"USDT_ETC":  true, // Ethereum Classic
	"USDT_ETH":  true, // Ethereum
	"USDT_LTC":  true, // Litecoin
	"USDT_NXT":  true, // NXT
	"USDT_REP":  true, // Augur
	"USDT_STR":  true, // Stellar
	"USDT_XMR":  true, // Monero
	"USDT_XRP":  true, // Ripple
	"USDT_ZEC":  true, // Zcash
}

var validPeriods = map[int]bool{
	300:   true,
	900:   true,
	1800:  true,
	7200:  true,
	14400: true,
	86400: true,
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, 500, "An error occurred.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func returnChartData(w http.ResponseWriter, r *http.Request, params url.Values) {
	currencyPair := params.Get("currencyPair")
	if !currencyPairs[currencyPair] {
		respondWithError(w, 400, "Invalid currencyPair")
		return
	}

	start, err := strconv.ParseInt(params.Get("start"), 10, 64) // 1405699200
	if err != nil || start < 0 {
		respondWithError(w, 400, "Invalid start")
		return
	}

	period64, err := strconv.ParseInt(params.Get("period"), 10, 16)
	if err != nil {
		respondWithError(w, 400, "Invalid period")
		return
	}
	period := int(period64)
	if !validPeriods[period] {
		respondWithError(w, 400, "Invalid period")
		return
	}

	end, err := strconv.ParseInt(params.Get("end"), 10, 64) // 9999999999
	if err != nil || end > 9999999999 || end < (start+period64) {
		respondWithError(w, 400, "Invalid end")
		return
	}

	len := int((end - start) / period64)

	quotes := make([]Quote, len)

	previousClose := rand.Float64() * 100

	for quote := 0; quote < len; quote++ {
		date := start + int64(quote*period)
		open := previousClose + ((rand.Float64() - 0.5) * 10)
		high := open + (rand.Float64() * 10)
		low := open + (rand.Float64() * -10)
		close := open + ((rand.Float64() - 0.5) * 20)
		volume := rand.Float64() * 100
		quoteVolume := rand.Float64() * 10000
		weightedAverage := close + ((rand.Float64() - 0.5) * 0.0001)
		previousClose = close

		quotes[quote] = Quote{
			Date:            date,
			Open:            open,
			High:            high,
			Low:             low,
			Close:           close,
			Volume:          volume,
			QuoteVolume:     quoteVolume,
			WeightedAverage: weightedAverage,
		}
	}

	respondWithJSON(w, 200, quotes)
}

func public(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	command := params.Get("command")
	if command != "returnChartData" {
		respondWithError(w, 501, "NOT IMPLEMENTED")
		return
	}

	returnChartData(w, r, params)
}

func main() {
	var realm string
	flag.StringVar(&realm, "realm", "realm1", "realm")

	var host string
	flag.StringVar(&host, "host", "", "host")

	var port int
	flag.IntVar(&port, "port", 8000, "port")

	var restAPIPort int
	flag.IntVar(&restAPIPort, "restAPIPort", 8001, "REST API port")

	flag.Parse()

	turnpikeAddr := fmt.Sprintf("%s:%d", host, port)
	restAPIAddr := fmt.Sprintf("%s:%d", host, restAPIPort)

	turnpike.Debug()

	s := turnpike.NewBasicWebsocketServer(realm)

	allowAllOrigin := func(r *http.Request) bool { return true }

	s.Upgrader.CheckOrigin = allowAllOrigin

	server := &http.Server{
		Addr:    turnpikeAddr,
		Handler: s,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/public", public)

	log.Println("turnpike server starting on", turnpikeAddr)
	go server.ListenAndServe()

	log.Println("rest api starting on", restAPIAddr)
	log.Fatal(http.ListenAndServe(restAPIAddr, mux))
}
