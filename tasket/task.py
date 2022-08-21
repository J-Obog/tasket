from dataclasses import dataclass
from dataclasses_json import DataClassJsonMixin
from datetime import datetime
from typing import Dict

@dataclass
class Task(DataClassJsonMixin):
    id: str
    body: Dict[str, any]
    exec_at: datetime
    sent_at: datetime