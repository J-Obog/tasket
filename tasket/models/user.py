from dataclasses import dataclass

@dataclass
class User:
    id: str
    email: str
    password: str
    created_at: int
    updated_at: int
