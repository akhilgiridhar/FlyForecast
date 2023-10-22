# FlyForecast
Created at HackTX 2023


Our web application uses ML to predict whether a particular flight will be delayed or not. We use custom REST APIs developed in Flask to interface with the model as well as APIs from weather sites to make predictions for a particular flight. In addition to the ML model, we also use Hedera for model provenance tracking. This creates a public ledger of the model's dataset, hyperparameters, training information, and creators, enabling greater transparency and ensuring that any bias or misuse of data can be seen by the public. We use React.js for the front end.

ML model based off of: [link to dataset]


## Initial Setup

1. Clone the repository: `git clone https://github.com/akhilgiridhar/FlyForecast.git`

## Dataset Setup

These instructions detail how to install the ASL Alphabet dataset. 
Other datasets can be used by creating a class which inherits from `torch.utils.data.Dataset`.

1. Install the [ASL Alphabet Dataset](https://www.kaggle.com/datasets/grassknoted/asl-alphabet) to `\data\`
2. Remove the `SPACE`, `DELETE`, and `NOTHING` folders from `\data\` as they are unused 
3. Convert the dataset to landmarks by calling `generateLandmarkDataset` in `train.py`

## Usage

`train.py` contains functions necessary to train new models.

`main.py` runs the selected model in real-time, taking video input from the primary webcam device.

## Authors
Created during the 2023 HackTX hackathon for team FlyForecast.
- Akhil Giridhar
- Rohain Jain
- Jibran Cutlerywala
- Boris He

