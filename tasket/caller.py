from datetime import datetime
from typing import Callable
from .ack import Ack
from .base import Client
import json

AckFunc = Callable[[Ack], None]

class Caller(Client):
    def __init__(self, url: str, channel: str, key: str, fn: AckFunc = None):
        super().__init__(url, "CALLER", channel, key)
        self._client.on("message", self._handle_ack)
        self._fn = fn

    def do(self, body: bytes, at: datetime):
        self._client.send(json.dumps({"body": body, "at": at}))

    def _handle_ack(self, payload: bytes) -> bytes:
        if self._fn:
            ack = Ack.from_json(payload)
            self._fn(ack) 
        
        return bytes("OK") 