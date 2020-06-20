from Unfurler import Unfurler
from flask import Flask, request, Response, jsonify

app = Flask(__name__)

@app.route("/")
def home():
    return {"msg": "Karibu Nyumbani"}

@app.route("/unfurl/", methods= ["POST"])
def unfurl():
    data = request.get_json()
    f = Unfurler.Unfurler(data.get('url'))
    data =f.unfurl()
    return jsonify(data)

if __name__ == "__main__":
    app.run(debug=True)