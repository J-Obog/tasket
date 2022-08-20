from typing import Callable
from .base import Client

class Handler(Client):
    def __init__(self, url: str, channel: str, key: str, callback: Callable):
        super().__init__(url, "HANDLER", channel, key)

    def _handle_ack(ack_msg: bytes):
        pass