from flask import Flask, request, jsonify
from Airline_Delay_Predictor import predictor

app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict():
    data = request.json
    prediction = predictor.get_prediction(data) 
    return jsonify({"prediction": prediction})

if __name__ == '__main__':
    app.run(port=5001)
