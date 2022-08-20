from datetime import datetime
from typing import Callable
from .base import Client

class Caller(Client):
    def __init__(self, url: str, channel: str, key: str, callback: Callable):
        super().__init__(url, "CALLER", channel, key)

    def do(body: bytes, at: datetime):
        pass