from dataclasses import dataclass
from datetime import datetime
from typing import Dict

@dataclass
class Ack:
    id: str
    task_id: str
    success: bool
    body: Dict[str, any]
    sent_at: datetime
    delta: int