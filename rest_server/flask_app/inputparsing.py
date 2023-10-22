import requests

# Define your API key
api_key = 'f69006857c004d90ad0222208232110'  # Replace with your actual API key

# Define the API endpoint you want to access
api_url = 'http://api.weatherapi.com/v1/forecast.json'  # Replace with the API's URL


def parsedata(city, time):
    #take the city and get the lat and long
    #take the time and use that to parse through the forecast data
    #make sure the time input is the 24-hour time and just the hour
    #eg: 7:45pm has an input of 17
    #dew, slp, tmp, vis, windspeed

    # Include the API key as a query parameter
    params = {'key': api_key, 'q': city}

    # Make the GET request
    response = requests.get(api_url, params=params)


    # Check if the request was successful
    if response.status_code == 200:
        data = response.json()  # Assuming the response is in JSON format
        # Process the data as needed
        Wind_speed = data['forecast']['forecastday'][0]['hour'][time]['wind_mph'] #hour array will change based off approximate time
        dewpoint = data['forecast']['forecastday'][0]['hour'][time]['dewpoint_f'] #hour array will change based off approximate ti
        slp = data['forecast']['forecastday'][0]['hour'][time]['pressure_mb'] #hour array will change based off approximate ti
        vis = data['forecast']['forecastday'][0]['hour'][time]['vis_km'] * 1000 #hour array will change based off approximate ti
        temperature_forecast = data['forecast']['forecastday'][0]['hour'][time]['temp_c'] #hour array will change based off approximate ti
        print(dewpoint, slp, temperature_forecast, vis, Wind_speed)
    else:
        print(f"Request failed with status code: {response.status_code}")

# if __name__ == '__main__':
#     parsedata('New York', 17)
