from datetime import datetime


class BaseComponent():
    def __init__(self) -> None:
        pass
    
    def create_times_step(self, start: datetime, end: datetime):
        time_dict = dict()
        time_dict['start'] = str(start.timestamp())
        time_dict['end'] = str(end.timestamp())
        return time_dict
    