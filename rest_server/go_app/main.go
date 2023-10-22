package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func getValidIntInput(prompt string) int {
	var value int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%d\n", &value)
		if err == nil {
			break
		} else {
			fmt.Println("Please enter a valid number.")
		}
	}
	return value
}

func main() {
	var input Input

	input.Dew = getValidIntInput("Dew: ")
	input.Slp = getValidIntInput("Slp: ")
	input.Tmp = getValidIntInput("Tmp: ")
	input.Vis = getValidIntInput("Vis: ")
	input.Wnd_speed = getValidIntInput("Wnd_speed: ")

	output := getPrediction(input)
	if output.Delay == 0 {
		fmt.Println("No Delay")
		return
	} else {
		fmt.Println("Delay")
	}
}
