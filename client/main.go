package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type WeatherResult struct {
	City        string  `json:"city"`
	Region      string  `json:"region"`
	Country     string  `json:"country"`
	LocalTime   string  `json:"local_time"`
	Temperature float64 `json:"temperature_c"`
	Humidity    int     `json:"humidity"`
	Condition   string  `json:"condition"`
	WindKph     float64 `json:"wind_kph"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter city (or type 'exit' to quit): ")
		city, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}

		city = strings.TrimSpace(city)
		if strings.ToLower(city) == "exit" || city == "" {
			fmt.Println("👋 Goodbye!")
			break
		}

		resp, err := http.Get("http://localhost:8080/weather?city=" + city)
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Server error:", resp.Status)
			continue
		}

		var w WeatherResult
		if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
			fmt.Println("Decode error:", err)
			continue
		}

		fmt.Printf("\n📍 Weather in %s, %s, %s:\n", w.City, w.Region, w.Country)
		fmt.Printf("   🕒 Local Time:   %s\n", w.LocalTime)
		fmt.Printf("   🌡️ Temperature:  %.1f °C\n", w.Temperature)
		fmt.Printf("   💧 Humidity:     %d%%\n", w.Humidity)
		fmt.Printf("   ☁️  Condition:    %s\n", w.Condition)
		fmt.Printf("   🍃 Wind:         %.1f kph\n\n", w.WindKph)
	}
}
