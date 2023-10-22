# Import necessary libraries
import xgboost as xgb
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score
from sklearn.linear_model import LogisticRegression
import joblib


def get_prediction(cig, dew, slp, tmp, vis, wind_speed):
    # Load the model
    modelXGB = joblib.load('xgb_trained_model.pkl')
    modelLogReg = joblib.load('LogReg_trained_model.pkl')

    # Create a DMatrix
    features = ['CIG', 'DEW', 'SLP', 'TMP', 'VIS', 'WND_SPD']
    data = [features, [cig, dew, slp, tmp, vis, wind_speed]]
    dmatrix = xgb.DMatrix(data)

    # Make a prediction
    xgb_prediction = modelXGB.predict(dmatrix)
    logreg_prediction = modelLogReg.predict(data)

    #print predictions of both models with probabilities
    print(f"XGB prediction: {xgb_prediction}")
    print(f"LogReg prediction: {logreg_prediction}")

    # Return the binary prediction based on majority vote
    prediction_binary = 1 if (xgb_prediction + logreg_prediction) >= 1 else 0
    return prediction_binary





    