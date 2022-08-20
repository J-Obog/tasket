from dataclasses import dataclass
from datetime import datetime
from typing import Dict

@dataclass
class Task:
    id: str
    body: Dict[str, any]
    exec_at: datetime
    sent_at: datetime