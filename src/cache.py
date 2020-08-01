from src import redis_

def check_url(url):
    is_url = redis_.get(url)
    if is_url is None: 
        print(url,"Not Found")
        raise KeyError
    print(url,"Found")
    return is_url

def add_url(url,data):
    print(url,"Added")
    redis_.set(url,data)
