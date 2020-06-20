from bs4 import BeautifulSoup
from typing import List, Dict
import requests
from pyoembed import oEmbed, PyOembedException

class Unfurler:

    def __init__(self,url: str):
        self.url: str = url
    
    def get_oembed_data(self):
        """Get the oembed data from server """

        try:
            oembed_data = oEmbed(self.url)
        except PyOembedException as e:
            print(e)
            return None
        else:
            o_list = [oembed_data]
            return {"oembed": o_list}
    
    def get_twitter_data(self, meta_tags: List):
        """
        Extracts the twitter meta Tags from the list of meta tags
        Adds the extracted twitter meta tags to the list of twitter tags *twitter_tags*

        Args:
            param1 meta_tags (List) - List containing meta tags
        """
        twitter_tags = []
        for meta in meta_tags:
            meta_name: str = meta.get("name")
            if meta_name and meta_name.startswith("twitter:"):
                split_tag: List = meta_name.split(":")
                if len(split_tag) > 2:# some tags have more than one colon separator e.g image:src
                    twit_tag_ = split_tag[1] + ":" + split_tag[2]
                else:
                    twit_tag_ = split_tag[1] 
                twitter_tags.append({twit_tag_:meta.get("content")})
        return {"twitter": twitter_tags}
    

    def get_og_tags(self, meta_tags : List):
        """
        Extracts the Open graph meta Tags from the list of meta tags
        Adds the extracted Open graph meta tags to the list of open_graph tags *og_tags*

        Args:
            param1 meta_tags (List) - List containing meta tags 
        """
        og_tags = []
        for meta in meta_tags:
            meta_name : str = meta.get("property")
            if meta_name and (meta_name.startswith("og:") or meta_name.startswith("fb:")):
                split_tag : List = meta_name.split(":")
                if len(split_tag) > 2:
                    twit_tag_ = split_tag[1] + ":" + split_tag[2]
                else:
                    twit_tag_ = split_tag[1] 
                og_tags.append({twit_tag_:meta.get("content")})
        return {"og": og_tags}
    
    def unfurl(self):
        """Entry point to class i.e. like run the class"""
        data = self.get_oembed_data()
        if data:
            return data
        else:
            resp = requests.get(self.url)
            soup = BeautifulSoup(resp.text, 'lxml')
            metas: List = soup.find_all("meta")
            tw_data = self.get_twitter_data(metas)
            og_data = self.get_og_tags(metas)
            if tw_data is None or og_data is None:
                return None
            if len(tw_data["twitter"]) >= len(og_data["og"]):
                return tw_data
            else:
                return og_data
        
        
    
