from . import Unfurler
from src.cache import check_url
import json
from flask import Flask, request, Response, jsonify, abort, Blueprint

app = Flask(__name__)
unfurler_bp = Blueprint("furler_bp",__name__,url_prefix="/unfurl")

@unfurler_bp.route("/home/")
def home():
    return {"msg": "Welcome to the unfurling site"}

@unfurler_bp.route("/unfurl/", methods= ["POST"])
def unfurl():
    data = request.get_json()
    url: str = data.get('url')
    if url is None:
        return abort(400,"Missing essentials")
    try:
        url_details=json.loads(check_url(url))
        return jsonify(url_details)
    except KeyError as e:
        f = Unfurler.Unfurler(url)
        data =f.unfurl()
        return jsonify(data)

