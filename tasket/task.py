from dataclasses import dataclass
from datetime import datetime

@dataclass
class Task:
    body: bytes
    exec_at: datetime
    timestamp: datetime