from dataclasses import dataclass
from datetime import datetime
from typing import Dict
from __future__ import annotations
import json

@dataclass
class Ack:
    id: str
    task_id: str
    success: bool
    body: Dict[str, any]
    sent_at: datetime
    delta: int

    def to_json(self) -> bytes:
        return bytes(json.dumps(self.__dict__), 'utf-8')

    @classmethod
    def from_json(cls, payload: bytes) -> Ack:
        obj = json.loads(payload)
        return Ack(**obj)