from dataclasses import dataclass
from datetime import datetime

@dataclass
class Ack:
    success: bool
    body: bytes
    timestamp: datetime
    delta: int