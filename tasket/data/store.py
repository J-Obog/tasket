from abc import ABC, abstractmethod
from tasket.data import Task, User, Log, TaskFilter, LogFilter
from typing import Optional, List

class UserStore(ABC):
    @abstractmethod
    def get(self, id: str) -> Optional[User]:
        ...

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
    def get(self, id: str, user_id: str) -> Optional[Task]:
        ...

    @abstractmethod
    def get_by(self, user_id: str, filter: TaskFilter) -> List[Task]:
        ...

    @abstractmethod
    def update(self, id: str, user_id: str, task: Task):
        ...

    @abstractmethod
    def insert(self, task: Task):
        ...


class LogStore(ABC):
    @abstractmethod
    def get_by(self, task_id: str, filter: LogFilter) -> List[Log]:
        ...