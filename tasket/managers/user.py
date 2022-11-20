from tasket.store import UserStore
from tasket.models import User, UserRequest
from tasket.utils import generate_uid, time_now, generate_hash
from typing import Optional


class UserManager:
    def __init__(self, user_store: UserStore):
        self._user_store = user_store

    def get_by_id(self, id: str) -> Optional[User]:
        return self._user_store.get_by_id(id)

    def create_user(self, user_req: UserRequest):
        id = generate_uid()
        curr_time = time_now()
        hashedpw = generate_hash(user_req.password)
        user = User(
            id=id,
            email=user_req.email,
            password=hashedpw,
            created_at=curr_time,
            updated_at=curr_time,
        )
        self._user_store.insert(user)
