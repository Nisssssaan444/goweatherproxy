// goweather - A simple weather  server in Go

package main

import (
    "encoding/json"
    "log"
    "net/http"
    "goweatherproxy/weather"
)

func main() {
    http.HandleFunc("/weather", weatherHandler)
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
    if city == "" {
        http.Error(w, "`city` param required", http.StatusBadRequest)
        return
    }

    data, err := weather.FetchWeather(city)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
    log.Printf("Fetched weather for %s", city)
}
