# Import necessary libraries
import xgboost as xgb
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score
from sklearn.linear_model import LogisticRegression
import joblib


# Load dataset (replace path with the actual dataset file path)
data = pd.read_csv("C:/Users/rohan/Downloads/airline_delay_dataset/Flight-weather-delay.csv")

# Assuming that the last column of your dataset is the target variable and the rest are features
X = data.iloc[:, :-1]  # Features
y = data.iloc[:, -1]   # Target variable

# Split the dataset into training and testing sets
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

def trainXGBmodel():
    # Create a DMatrix (XGBoost's internal data structure) for training and testing data
    dtrain = xgb.DMatrix(X_train, label=y_train)
    dtest = xgb.DMatrix(X_test, label=y_test)
    print(dtest.feature_names)

    # Define the parameters for the XGBoost model
    params = {
        'objective': 'binary:logistic',  # For binary classification
        'max_depth': 3,                  # Maximum depth of trees
        'eta': 0.1,                      # Learning rate
        'eval_metric': 'logloss'         # Evaluation metric
    }

    # Train the XGBoost model
    num_round = 100  # Number of boosting rounds
    model = xgb.train(params, dtrain, num_round)

    # Make predictions on the test data
    y_pred = model.predict(dtest)

    # Convert the probabilities to binary predictions
    y_pred_binary = [1 if p > 0.5 else 0 for p in y_pred]

    # Calculate and print the accuracy
    accuracy = accuracy_score(y_test, y_pred_binary)
    print(f"XGB accuracy: {accuracy * 100:.2f}%")

    # Serialize and save the model
    model_filename = 'xgb_trained_model.pkl'  # Choose a file name and extension
    joblib.dump(model, model_filename)

def trainLogRegModel():
    # Create a logistic regression model
    model = LogisticRegression()

    # Train the model on the training data
    model.fit(X_train, y_train)

    # Make predictions on the test data
    y_pred = model.predict(X_test)

    # Calculate and print the accuracy
    accuracy = accuracy_score(y_test, y_pred)
    print(f"Logistical Regression accuracy: {accuracy * 100:.2f}%")

    # Serialize and save the model
    model_filename = 'LogReg_trained_model.pkl'  # Choose a file name and extension
    joblib.dump(model, model_filename)

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

#write main function that calls the above functions
if __name__ == '__main__':
    # trainXGBmodel()
    # trainLogRegModel()
    
    cig = 2
    get_prediction(22000, -117, 10214, -44, 16093, 46)





    