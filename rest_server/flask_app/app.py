from flask import Flask, request, jsonify
from Airline_Delay_Predictor import predictor

app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict():
    data = request.json
    prediction = predictor.get_prediction(data['Cig'], data['Dew'], data['Slp'], data['Tmp'], data['Vis'], data['Wnd_speed'])     
    return jsonify({"prediction": prediction})

if __name__ == '__main__':
    app.run(port=5001)
