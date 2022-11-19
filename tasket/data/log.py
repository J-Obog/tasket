from dataclasses import dataclass
from enum import IntEnum

class LogSource(IntEnum):
    UNKNOWN = 0
    TASK = 1
    SERVER = 2

@dataclass
class Log:
    id: str
    task_id: str
    source: LogSource
    content:bytes
    created_at: int
