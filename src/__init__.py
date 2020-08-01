import os
from flask import Flask
from flask_redis import FlaskRedis
redis_ = FlaskRedis()
def create_app():
    main_app = Flask(__name__)
    main_app.config['DEBUG']=True
    print(os.environ.get("REDIS_URL"))
    main_app.config['REDIS_URL']=os.environ.get("REDIS_URL")
    with main_app.app_context():
        redis_.init_app(main_app)
        from src.Unfurler import unfurler_bp
        main_app.register_blueprint(unfurler_bp)
    return main_app