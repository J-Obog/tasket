from dataclasses import dataclass


@dataclass
class User:
    id: str
    email: str
    password: bytes
    created_at: int
    updated_at: int


@dataclass
class UserRequest:
    email: str
    password: str
