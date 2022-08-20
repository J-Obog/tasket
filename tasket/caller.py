from datetime import datetime
from typing import Callable
import socketio

class Caller():
    def __init__(self, url: str, channel: str, key: str, callback: Callable):
        headers = {"authorization": key, "client-type": "CALLER", "channel": channel}
        self._client = socketio.Client()
        self._client.connect(url, headers)

    def do(body: bytes, at: datetime):
        pass

    def close(self):
        self._client.disconnect()
        
    def _handle_ack(ack_msg: bytes):
        pass



