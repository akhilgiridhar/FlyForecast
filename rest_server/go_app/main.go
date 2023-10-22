package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Input struct {
	Dew       int `json:"dew"`
	Slp       int `json:"slp"`
	Tmp       int `json:"tmp"`
	Vis       int `json:"vis"`
	Wnd_speed int `json:"wnd_speed"`
}

type Output struct {
	Delay int `json:"delay"`
}

const api_key = "f69006857c004d90ad0222208232110"            // Replace with your actual API key
const api_url = "http://api.weatherapi.com/v1/forecast.json" // Replace with the API's URL

func callParsedata(city, time string) Input {
	var result Input

	// Create URL with parameters
	u, _ := url.Parse(api_url)
	q := u.Query()
	q.Set("key", api_key)
	q.Set("q", city)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Failed to make request:", err)
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return result
	}

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	forecast := data["forecast"].(map[string]interface{})
	day := forecast["forecastday"].([]interface{})[0]
	hourData := day.(map[string]interface{})["hour"].([]interface{})

	timeInt, _ := strconv.Atoi(time)
	hour := hourData[timeInt].(map[string]interface{})

	result.Dew = int(hour["dewpoint_f"].(float64))
	result.Slp = int(hour["pressure_mb"].(float64))
	result.Tmp = int(hour["temp_c"].(float64))
	result.Vis = int(hour["vis_km"].(float64) * 1000)
	result.Wnd_speed = int(hour["wind_mph"].(float64))

	return result
}

func getPrediction(data Input) Output {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling data: %v", err)
	}

	resp, err := http.Post("http://localhost:5001/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response map[string]int
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	return Output{Delay: response["prediction"]}
}

func main() {
	var input Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter City: ")
	city, _ := reader.ReadString('\n')
	city = city[:len(city)-1] // Remove the trailing newline

	fmt.Print("Enter Time: ")
	time, _ := reader.ReadString('\n')
	time = time[:len(time)-1] // Remove the trailing newline

	fmt.Println("City:", city)
	fmt.Println("Time:", time)

	input = callParsedata(city, time)
	fmt.Println(input)

	output := getPrediction(input)
	if output.Delay == 0 {
		fmt.Println("No Delay")
		return
	} else {
		fmt.Println("Delay")
	}
}
