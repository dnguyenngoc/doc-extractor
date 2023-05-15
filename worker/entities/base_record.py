from faust import Record


class BaseRecord(Record):
    request_id: str
    request_type: str
    pipeline: dict
    data: dict


