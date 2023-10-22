import numpy as np
import pandas as pd
import xgboost as xgb
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score



def get_prediction(dew, slp, tmp, vis, wind_speed):
    data = pd.read_csv("Flight-weather-delay-correlation-data.csv")

    X = data.drop('WEATHER_DELAY', axis=1)
    y = data['WEATHER_DELAY']

    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    model = xgb.XGBClassifier()

    model.fit(X_train, y_train)

    y_pred = model.predict(X_test)

    accuracy = accuracy_score(y_test, y_pred)
    print(f"Accuracy: {accuracy * 100:.2f}%")

    sample = np.array([[dew, slp, tmp, vis, wind_speed]])
    prediction = model.predict(sample)
    return prediction[0]

