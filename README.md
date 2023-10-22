# FlyForecast
Created at HackTX 2023

Made by:
Rohan Jain
Akhil Giridhar
Jibran Cutlerywala
Boris He

Our web application uses ML to predict whether a particular flight will be delayed or not. We use custom REST APIs developed in Flask to interface with the model as well as APIs from weather sites to make predictions for a particular flight. In addition to the ML model, we also use Hedera for model provenance tracking. This creates a public ledger of the model's dataset, hyperparameters, training information, and creators, enabling greater transparency and ensuring that any bias or misuse of data can be seen by the public. We use React.js for the front end.

ML model based off of: [link to dataset]
