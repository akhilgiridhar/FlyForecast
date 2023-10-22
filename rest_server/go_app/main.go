package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

type Input struct {
	Dew       int `json:"dew"`
	Slp       int `json:"slp"`
	Tmp       int `json:"tmp"`
	Vis       int `json:"vis"`
	Wnd_speed int `json:"wnd_speed"`
}

type api_params struct {
	City string `json:"city"`
	Time int    `json:"time"`
}

type Output struct {
	Delay int `json:"delay"`
}

func callParsedata(city, time string) Input {
	cmd := exec.Command("python", "inputparsing.py", "parsedata", time, city)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		return Input{}
	}

	// The output from the Python script is stored in 'out'
	// Unmarshal the JSON data into the Go struct
	pythonOutput := out.String()

	// Unmarshal the JSON data into the Go struct
	var input Input
	err = json.Unmarshal([]byte(pythonOutput), &input)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return Input{}
	}
	return input
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

func getValidIntInput(prompt string) string {
	var value string
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%d\n", &value)
		if err == nil {
			break
		} else {
			fmt.Println("Please enter a valid string.")
		}
	}
	return value
}

func main() {
	var input Input
	// var api_params Params
	var city, time string
	city = getValidIntInput("Enter City: ")
	time = getValidIntInput("Enter Time: ")

	input = callParsedata(city, time)

	output := getPrediction(input)
	if output.Delay == 0 {
		fmt.Println("No Delay")
		return
	} else {
		fmt.Println("Delay")
	}
}
