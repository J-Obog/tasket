from abc import ABC, abstractmethod
from tasket.models import (
    Task,
    TaskFilter,
    TaskStatus,
    User,
    Log,
    LogFilter,
)

from typing import Optional, List


class UserStore(ABC):
    @abstractmethod
    def get_by_id(self, id: str) -> Optional[User]:
        ...

    # TODO: handle different update scenarios
    @abstractmethod
    def update(self, id: str, user: User):
        ...

    @abstractmethod
    def insert(self, user: User):
        ...

    @abstractmethod
    def delete(self, id: str):
        ...


class TaskStore(ABC):
    @abstractmethod
    def get_by_id(self, id: str, user_id: str) -> Optional[Task]:
        ...

    @abstractmethod
    def get_by_user(self, user_id: str) -> List[Task]:
        ...

    @abstractmethod
    def get_by_filters(self, user_id: str, filter: TaskFilter) -> List[Task]:
        ...

    @abstractmethod
    def update_name(self, id: str, user_id: str, name: str):
        ...

    @abstractmethod
    def update_status(self, id: str, status: TaskStatus):
        ...

    @abstractmethod
    def insert(self, task: Task):
        ...


class LogStore(ABC):
    @abstractmethod
    def get_by_task(self, task_id: str) -> List[Log]:
        ...

    @abstractmethod
    def get_by_filters(self, task_id: str, filter: LogFilter) -> List[Log]:
        ...

    @abstractmethod
    def insert(self, log: Log):
        ...
