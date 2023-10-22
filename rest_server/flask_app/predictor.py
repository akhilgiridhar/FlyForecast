# Import necessary libraries
import xgboost as xgb
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score, classification_report
from sklearn.linear_model import LogisticRegression
import joblib


def get_prediction(dew, slp, tmp, vis, wind_speed):
    

    # 1. Load the data
    data = pd.read_csv('path_to_your_data.csv')

    # 2. Data Pre-processing
    # (Assuming data is already clean and does not have missing values.)

    # 3. Split the data
    X = data[['dew', 'slp', 'tmp', 'vis', 'wind_speed']]
    y = data['weather_delay']  # Assuming 1 for delay, 0 for no delay

    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # 4. Model Training
    model = xgb.XGBClassifier()
    model.fit(X_train, y_train)

    # 5. Model Evaluation
    y_pred = model.predict(X_test)

    accuracy = accuracy_score(y_test, y_pred)
    print(f"Accuracy: {accuracy:.2f}")
    print(classification_report(y_test, y_pred))

    # For future predictions:
    # new_data = pd.DataFrame({'dew': [value], 'slp': [value], 'tmp': [value], 'vis': [value], 'wind_speed': [value]})
    # prediction = model.predict(new_data)
    # print(prediction)

if __name__ == '__main__':
    get_prediction(1, 2, 3, 4, 5)




    