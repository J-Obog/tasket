from typing import Callable
import socketio

class Handler:
    def __init__(self, url: str, channel: str, key: str, callback: Callable):
        headers = {"authorization": key, "client-type": "HANDLER", "channel": channel}
        self._client = socketio.Client()
        self._client.connect(url, headers)

    def close(self):
        self._disconnect()

    def _handle_ack(ack_msg: bytes):
        pass