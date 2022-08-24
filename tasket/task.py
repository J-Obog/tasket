from dataclasses import dataclass
from dataclasses_json import DataClassJsonMixin
from datetime import datetime

@dataclass
class Task(DataClassJsonMixin):
    id: str
    body: bytes
    exec_at: datetime
    sent_at: datetime