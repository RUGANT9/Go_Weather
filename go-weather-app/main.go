package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const apiKey = "56d1172146caccace958f8e7a39cb0fa"

func getWeather(city string) {
    url := fmt.Sprintf(
        "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
        city, apiKey,
    )

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Println("Failed to fetch weather. Status:", resp.Status)
        return
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

	if resp.StatusCode != 200 {
        // Show custom error message from response body if available
        if msg, ok := result["message"].(string); ok {
            fmt.Printf("❌ Failed to fetch weather: %s\n", msg)
        } else {
            fmt.Printf("❌ Failed to fetch weather. Status code: %d\n", resp.StatusCode)
        }
        return
    }

	main := result["main"].(map[string]interface{})
    weatherArr := result["weather"].([]interface{})
    weather := weatherArr[0].(map[string]interface{})

    temp := main["temp"]
    humidity := main["humidity"]
    description := weather["description"]

    fmt.Println("------ Weather Report ------")
    fmt.Printf("City: %s\n", city)
    fmt.Printf("Temperature: %.1f°C\n", temp)
    fmt.Printf("Humidity: %v%%\n", humidity)
    fmt.Printf("Condition: %s\n", description)
}

func main() {
    var city string
    fmt.Print("Enter a city: ")
    fmt.Scanln(&city)

    getWeather(city)
}
