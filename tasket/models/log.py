from dataclasses import dataclass
from enum import IntEnum
from typing import Optional

class LogSource(IntEnum):
    UNKNOWN = 0
    TASK = 1
    SERVER = 2

@dataclass
class Log:
    id: str
    task_id: str
    source: LogSource
    content: bytes
    created_at: int


@dataclass
class LogFilter:
    source: Optional[LogSource] = None
    contains: Optional[bytes] = None
    created_after: Optional[int] = None
    created_before: Optional[int] = None
