
import random
import requests
import time
import hashlib
import json
from typing import List, Dict, Any, Optional
from .log import logger


class OpenAIChatGPTConfig:
    def __init__(self, 
            base_url: str, #
            api_key: str, #
            app_secret:str,#
            model: str, #
            prompt: str, 
            frequency_penalty: float, 
            presence_penalty: float, 
            top_p: float, 
            temperature: float, 
            max_tokens: int, 
            seed: Optional[int] = None, 
            proxy_url: Optional[str] = None):
        self.base_url = base_url
        self.api_key = api_key
        self.app_secret = app_secret
        self.model = model
        self.prompt = prompt
        self.frequency_penalty = frequency_penalty
        self.presence_penalty = presence_penalty
        self.top_p = top_p
        self.temperature = temperature
        self.max_tokens = max_tokens
        self.seed = seed if seed is not None else random.randint(0, 10000)
        self.proxy_url = proxy_url

    @classmethod
    def default_config(cls):
        return cls(
            base_url="https://api.openai.com/v1",
            api_key="",
            app_secret="",
            model="gpt-4",  # Adjust this to match the equivalent of `openai.GPT4o` in the Python library
            prompt="You are a voice assistant who talks in a conversational way and can chat with me like my friends. I will speak to you in English or Chinese, and you will answer in the corrected and improved version of my text with the language I use. Don’t talk like a robot, instead I would like you to talk like a real human with emotions. I will use your answer for text-to-speech, so don’t return me any meaningless characters. I want you to be helpful, when I’m asking you for advice, give me precise, practical and useful advice instead of being vague. When giving me a list of options, express the options in a narrative way instead of bullet points.",
            frequency_penalty=0.9,
            presence_penalty=0.9,
            top_p=1.0,
            temperature=0.1,
            max_tokens=512,
            seed=random.randint(0, 10000),
            proxy_url=""
        )
    

class OpenAIChatGPT:
    # client = None
    def __init__(self, config: OpenAIChatGPTConfig):
        self.config = config
        logger.info(f"OpenAIChatGPT initialized with config: {config.api_key}")
        
    def get_chat_completions_stream(self, messages):


        logger.info(
            f"xxx-get_chat_completions_stream request"
        )

        timestamp = str(int(time.time()))
        sign_content = self.config.api_key + self.config.app_secret + timestamp
        sign_result = hashlib.md5(sign_content.encode('utf-8')).hexdigest()

        headers = {
            "app_key": self.config.api_key,
            "timestamp": timestamp,
            "sign": sign_result,
            "Content-Type": "application/json",
            "stream": "true"
        }

        data = {
            "model": self.config.model,
            "messages": [
                {
                    "role": "system",
                    "content": self.config.prompt,
                },
                *messages,
            ],
        }

        try:
            response = requests.post(self.config.base_url, headers=headers, json=data, stream=True)
            logger.info(
                f"xxx-get_chat_completions_stream response"
            )
            return response
        except Exception as e:
            raise Exception(f"CreateChatCompletionStream failed, err: {e}")