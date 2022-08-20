from dataclasses import dataclass
from datetime import datetime
from typing import Dict
from __future__ import annotations
import json

@dataclass
class Task:
    id: str
    body: Dict[str, any]
    exec_at: datetime
    timestamp: datetime

    def to_json(self) -> bytes:
        return bytes(json.dumps(self.__dict__), 'utf-8')

    @classmethod
    def from_json(cls, payload: bytes) -> Task:
        obj = json.loads(payload)
        return Task(**obj)