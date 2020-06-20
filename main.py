# Unfurling looks for following tags : osEmbed, twitter("twitter:"), opengraph("og:"), html meta-tags
from bs4 import BeautifulSoup
from typing import List
import requests

twitter_tags = []
og_tags =[]
def get_twitter_tags(meta_tags : List ):
    #Twitter uses meta["name"]
    for meta in meta_tags:
        meta_name:str = meta.get("name")
        if meta_name and meta_name.startswith("twitter:"):
            split_tag : List = meta_name.split(":")
            if len(split_tag) > 2:
                twit_tag_ = split_tag[1] + ":" + split_tag[2]
            else:
                twit_tag_ = split_tag[1] 
            twitter_tags.append({twit_tag_:meta.get("content")})


response = requests.Response()
domain = "neilPatel"
id = "OpenGraph"
file_name = "./saved/"+domain+id+".html"
url = "https://neilpatel.com/blog/open-graph-meta-tags/"
try:
    file = open(file_name)
    resp_text = file.read()
    file.close
except FileNotFoundError:
    print("Making request to server")
    response = requests.get(url)
    resp_text = response.text
    with open(file_name,"w") as git:
        git.write(resp_text)

soup = BeautifulSoup(resp_text,"lxml")
metas = soup.find_all('meta')
get_twitter_tags(metas)
for t_tag in twitter_tags:
    print(t_tag)