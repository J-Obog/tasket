from dataclasses import dataclass
from typing import Dict, Any


@dataclass
class RestRequest:
    url: str
    url_params: Dict[str, str]
    query_params: Dict[str, str]
    body: Dict[str, Any]
    headers: Dict[str, str]


@dataclass
class RestResponse:
    data: str
    status: int
