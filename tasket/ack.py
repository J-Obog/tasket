from dataclasses import dataclass
from dataclasses_json import DataClassJsonMixin
from datetime import datetime
from typing import Dict

@dataclass
class Ack(DataClassJsonMixin):
    id: str
    task_id: str
    success: bool
    body: bytes
    sent_at: datetime
    delta: int