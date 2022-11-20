from tasket.store import TaskStore
from tasket.models import Task, TaskRequest, TaskFilter, TaskStatus
from typing import Optional, List
from tasket.utils import generate_uid, time_now


class TaskManager:
    def __init__(self, task_store: TaskStore):
        self._task_store = task_store

    def get_by_id(self, id: str, user_id: str) -> Optional[Task]:
        return self._task_store.get_by_id(id, user_id)

    def get_by_user(self, user_id: str) -> List[Task]:
        return self._task_store.get_by_user(user_id)

    def get_by_filters(self, user_id: str, filter: TaskFilter) -> List[Task]:
        return self._task_store.get_by_filters(user_id, filter)

    def create_task(self, task_req: TaskRequest):
        id = generate_uid()
        curr_time = time_now()

        task = Task(
            id=id,
            user_id=task_req.user_id,
            name=task_req.name,
            status=TaskStatus.PENDING,
            max_timeout=task_req.max_timeout,
            max_memory=task_req.max_memory,
            commands=task_req.commands,
            environ=task_req.environ,
            sourceMeta=task_req.sourceMeta,
            started_at=None,
            completed_at=None,
            created_at=curr_time,
            updated_at=curr_time,
        )

        self._task_store.insert(task)

    def update_name(self, id: str, user_id: str, name: str):
        self._task_store.update_name(id, user_id, name)

    def update_status(self, id: str, status: TaskStatus):
        self._task_store.update_status(id, status)

    def stop_task(self, id: str, user_id: str):
        pass
