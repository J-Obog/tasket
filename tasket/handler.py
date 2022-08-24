from typing import Callable, Tuple
from .base import Client
from .task import Task
import json

TaskFunc = Callable[[Task], Tuple[bytes, bool]]

class Handler(Client):
    def __init__(self, url: str, channel: str, key: str = "", fn: TaskFunc = None):
        super().__init__(url, "HANDLER", channel, key)
        self._client.on("message", self._handle_task)
        self._fn = fn

    def _handle_task(self, payload: bytes) -> bytes:
        task = Task.from_json(payload)
        if self._fn:
            res, success = self._fn(task)
            return bytes(json.dumps({"res": res, "success": success}))
        
        return bytes('{"res":"", "success":"true"}')