from dataclasses import dataclass
from typing import List, Dict, Optional
from enum import IntEnum

class TaskStatus(IntEnum):
    UNKNOWN = 0
    COMPLETED = 1
    STOPPED = 2
    RUNNING = 3
    PENDING = 4

@dataclass
class Task:
    id: str
    name: str
    status: TaskStatus
    max_timeout: int
    max_memory: int
    commands: List[str]
    environVars: Dict[str,str]
    sourceMeta: Dict[str,str]
    started_at: Optional[int]
    completed_at: Optional[int]
    created_at: int
    updated_at: int
