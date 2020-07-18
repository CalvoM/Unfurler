from Unfurler import Unfurler
from flask import Flask, request, Response, jsonify, abort

app = Flask(__name__)

@app.route("/")
def home():
    return {"msg": "Welcome to the unfurling site"}

@app.route("/unfurl/", methods= ["POST"])
def unfurl():
    data = request.get_json()
    url: str = data.get('url')
    if url is None:
        return abort(400,"Missing essentials")
    f = Unfurler.Unfurler(url)
    data =f.unfurl()
    return jsonify(data)

if __name__ == "__main__":
    app.run()