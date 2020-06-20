# Unfurling looks for following tags : osEmbed, twitter("twitter:"), opengraph("og:"), html meta-tags
from bs4 import BeautifulSoup
from typing import List
import requests


def get_twitter_tags(meta_tags: List, twitter_tags: List):
    """
    Extracts the twitter meta Tags from the list of meta tags
    Adds the extracted twitter meta tags to the list of twitter tags *twitter_tags*

    Args:
        param1 meta_tags (List) - List containing meta tags
        param2 twitter_tags (List) - List to hold twitter meta tags
    """
    for meta in meta_tags:
        meta_name: str = meta.get("name")
        if meta_name and meta_name.startswith("twitter:"):
            split_tag: List = meta_name.split(":")
            if len(split_tag) > 2:# some tags have more than one colon separator e.g image:src
                twit_tag_ = split_tag[1] + ":" + split_tag[2]
            else:
                twit_tag_ = split_tag[1] 
            twitter_tags.append({twit_tag_:meta.get("content")})

def get_og_tags(meta_tags : List, og_tags : List):
    """
    Extracts the Open graph meta Tags from the list of meta tags
    Adds the extracted Open graph meta tags to the list of open_graph tags *og_tags*

    Args:
        param1 meta_tags (List) - List containing meta tags 
        param2 og_tags (List) - List to hold og meta tags
    """
    for meta in meta_tags:
        meta_name : str = meta.get("property")
        if meta_name and (meta_name.startswith("og:") or meta_name.startswith("fb:")):
            split_tag : List = meta_name.split(":")
            if len(split_tag) > 2:
                twit_tag_ = split_tag[1] + ":" + split_tag[2]
            else:
                twit_tag_ = split_tag[1] 
            og_tags.append({twit_tag_:meta.get("content")})

if __name__ == "__main__":
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
    og_tags =[]
    get_og_tags(metas, og_tags)
    for t_tag in og_tags:
        print(t_tag)
