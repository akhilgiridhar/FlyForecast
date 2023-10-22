from flask import Flask, request, jsonify
import predictor

app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict():
    data = request.json
    # print(data)
    # prediction = predictor.get_prediction(data['dew'], data['slp'], data['tmp'], data['vis'], data['wnd_speed'])     
    # return jsonify({"prediction": prediction})

if __name__ == '__main__':
    app.run(port=5001)
