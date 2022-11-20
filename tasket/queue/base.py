from abc import ABC, abstractmethod
from dataclasses import dataclass


@dataclass
class Message:
    id: str
    data: bytes
    timestamp: int


class Queue:
    @abstractmethod
    def push(self, message: Message):
        ...

    @abstractmethod
    def pull(self) -> Message:
        ...
