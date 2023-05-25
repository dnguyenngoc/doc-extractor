from enum import Enum, auto

class RequestStatus(Enum):
    TIME_OUT = auto()
    PENDING = auto()
    SUCCESS = auto()
    FAILED = auto()