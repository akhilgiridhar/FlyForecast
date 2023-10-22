# FlyForecast
Created at HackTX 2023


Winner of best Use of Hedera

Our web application uses ML to predict whether a particular flight will be delayed or not. We use custom REST APIs developed in Flask to interface with the model as well as APIs from weather sites to make predictions for a particular flight. In addition to the ML model, we also use Hedera for model provenance tracking. This creates a public ledger of the model's dataset, hyperparameters, training information, and creators, enabling greater transparency and ensuring that any bias or misuse of data can be seen by the public. We use React.js for the front end.

## Initial Setup

1. Clone the repository: `git clone https://github.com/akhilgiridhar/FlyForecast.git`

## Dataset Setup

These instructions detail how to install the ASL Alphabet dataset. 
Other datasets can be used by creating a class which inherits from `torch.utils.data.Dataset`.

1. Install the [dataset](https://github.com/nitilaksha1/Analysis-of-Flight-Delay-and-Weather-Dataset/blob/master/machine-learning/Flight-weather-delay-correlation-data.csv) to `\data\`
2. Remove the `FLIGHT_ID` column and move the  `WEATHER_DELAY` column to the end from `\data\` as they are unused 

## Usage

`predictor.py` contains functions necessary to train the model.

`main.py` runs the selected model in real-time, taking inputs for location and time from the UI built in React.js.

## Authors
Created during the 2023 HackTX hackathon for team FlyForecast.
- Akhil Giridhar
- Rohan Jain
- Jibran Cutlerywala
- Boris He

