from tasket.store import LogStore
from tasket.models import Log, LogFilter, LogRequest
from typing import List
from tasket.utils import generate_uid


class LogManager:
    def __init__(self, log_store: LogStore):
        self._log_store = log_store

    def get_by_task(self, task_id: str) -> List[Log]:
        return self._log_store.get_by_task(task_id)

    def get_by_filters(self, task_id: str, filter: LogFilter) -> List[Log]:
        return self._log_store.get_by_filters(task_id, filter)

    def create_log(self, log_req: LogRequest):
        id = generate_uid()
        log = Log(
            id=id,
            task_id=log_req.task_id,
            source=log_req.source,
            content=log_req.content,
            created_at=log_req.created_at,
        )
        self._log_store.insert(log)
