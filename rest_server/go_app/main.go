package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
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

	output := getPrediction(input)
	if output.Delay == 0 {
		fmt.Println("No Delay")
		return
	} else {
		fmt.Println("Delay")
	}
}
