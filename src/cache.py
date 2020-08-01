from src import redis_

def check_url(url):
    is_url = redis_.get(url)
    if is_url is None: 
        raise KeyError
    return is_url

def add_url(url,data):
    redis_.set(url,data)
