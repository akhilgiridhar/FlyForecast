package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"net"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"rest_server/go_app/utils"

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

func print_output(prediction int) {
	if prediction == 0 {
		fmt.Println("Your flight has a high chance of being on time")
	} else {
		fmt.Println("Your flight has a high chance of being delayed")
	}
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

func handlePredictDelay(w http.ResponseWriter, r *http.Request) {
	// Decode the input data from request body
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Bad request data", http.StatusBadRequest)
		return
	}
	city, ok1 := requestData["city"]
	time, ok2 := requestData["time"]

	if !ok1 || !ok2 {
		http.Error(w, "Missing city or time", http.StatusBadRequest)
		return
	}

	// Get the parsed weather data
	parsedData := callParsedata(city, time)

	// Predict using the weather data
	predictionOutput := getPrediction(parsedData)

	// Return the prediction
	responseData := map[string]int{
		"Delay": predictionOutput.Delay,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func getIP() string{
	// Get the host's IP addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return("Error:") + err.Error()
	}

	// Find the first non-loopback IP address
	var hostIP string
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			hostIP = ipNet.IP.String()
			break
		}
	}

	if hostIP != "" {
		return hostIP
	} else {
		return ("Unable to determine the host's IP address.")
	}
}

func main() {
	// var input Input
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter City: ")
	// city, _ := reader.ReadString('\n')
	// city = city[:len(city)-1] // Remove the trailing newline

	// fmt.Print("Enter Time: ")
	// time, _ := reader.ReadString('\n')
	// time = time[:len(time)-1] // Remove the trailing newline

	// input = callParsedata(city, time)
	// output := getPrediction(input)
	// print_output(output.Delay)
	client := utils.SetupHedera()
	var topicID = utils.GetTopicInfo("0.0.5739611")
	utils.SubscribeToTopic(client, topicID)

	r := mux.NewRouter()
	r.HandleFunc("/predict-delay", handlePredictDelay).Methods("POST")

	// Get the current time
	currentTime := time.Now()

	// Format the time as a string
	timeString := currentTime.Format("2006-01-02 15:04:05")
	transmitString := "Accessing predictor.py at " + timeString + " from IP address " + getIP()
	utils.TransmitMessage(transmitString, client, topicID)

	// Handle CORS
	handler := cors.Default().Handler(r)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
