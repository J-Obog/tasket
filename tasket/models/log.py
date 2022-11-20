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
class LogRequest:
    task_id: str
    source: LogSource
    content: bytes
    created_at: int


@dataclass
class LogFilter:
    source: Optional[LogSource]
    contains: Optional[bytes]
    created_after: Optional[int]
    created_before: Optional[int]
    first: int
    last: int
